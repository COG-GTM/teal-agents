import asyncio
import threading
from types import SimpleNamespace

import pytest
from pynamodb.models import DoesNotExist

from data.impl import dynamo_chat_history_manager as module
from data.impl.dynamo_chat_history_manager import DynamoChatHistoryManager
from model import ChatHistoryItem, MessageType


def _manager() -> DynamoChatHistoryManager:
    return object.__new__(DynamoChatHistoryManager)


ITEMS = [
    ChatHistoryItem(timestamp=1.0, agent_name="a", message_type=MessageType.USER, message="hi")
]


def test_get_chat_history_session_runs_lookups_concurrently(monkeypatch):
    both_started = threading.Barrier(2, timeout=5)

    async def fake_items(self, orchestrator_name, session_id):
        await asyncio.to_thread(both_started.wait)
        return ITEMS

    def fake_get(user_id, session_id):
        both_started.wait()
        return SimpleNamespace(orchestrator="ao", previous_session="prev")

    monkeypatch.setattr(DynamoChatHistoryManager, "get_session_items", fake_items)
    monkeypatch.setattr(module.DynamoChatHistory, "get", fake_get)

    history = asyncio.run(_manager().get_chat_history_session("ao", "u", "s"))

    assert history.user_id == "u"
    assert history.session_id == "s"
    assert history.previous_session == "prev"
    assert history.history == ITEMS


def test_get_chat_history_session_other_orchestrator_returns_none(monkeypatch):
    async def fake_items(self, orchestrator_name, session_id):
        return ITEMS

    monkeypatch.setattr(DynamoChatHistoryManager, "get_session_items", fake_items)
    monkeypatch.setattr(
        module.DynamoChatHistory,
        "get",
        lambda u, s: SimpleNamespace(orchestrator="other", previous_session=None),
    )

    assert asyncio.run(_manager().get_chat_history_session("ao", "u", "s")) is None


def test_get_chat_history_session_missing_session_raises(monkeypatch):
    async def fake_items(self, orchestrator_name, session_id):
        return ITEMS

    def fake_get(user_id, session_id):
        raise DoesNotExist()

    monkeypatch.setattr(DynamoChatHistoryManager, "get_session_items", fake_items)
    monkeypatch.setattr(module.DynamoChatHistory, "get", fake_get)

    with pytest.raises(DoesNotExist):
        asyncio.run(_manager().get_chat_history_session("ao", "u", "s"))


def test_get_chat_history_session_item_error_takes_precedence(monkeypatch):
    async def fake_items(self, orchestrator_name, session_id):
        raise RuntimeError("query failed")

    def fake_get(user_id, session_id):
        raise DoesNotExist()

    monkeypatch.setattr(DynamoChatHistoryManager, "get_session_items", fake_items)
    monkeypatch.setattr(module.DynamoChatHistory, "get", fake_get)

    with pytest.raises(RuntimeError, match="query failed"):
        asyncio.run(_manager().get_chat_history_session("ao", "u", "s"))
