"""External API mocking utilities for E2E tests.

This module provides mock server functionality for external APIs
like universities.hipolabs.com and Gemini API during testing.

TODO: Implement in Phase 2
"""

from typing import Any


class MockServer:
    """Mock server for external API responses."""

    def __init__(self, port: int = 8888):
        """Initialize the mock server.

        Args:
            port: Port number for the mock server
        """
        self.port = port
        self.server = None
        self.routes: dict[str, Any] = {}

    def start(self) -> bool:
        """Start the mock server.

        Returns:
            True if server started successfully, False otherwise
        """
        raise NotImplementedError("To be implemented in Phase 2")

    def stop(self) -> bool:
        """Stop the mock server.

        Returns:
            True if server stopped successfully, False otherwise
        """
        raise NotImplementedError("To be implemented in Phase 2")

    def add_route(self, path: str, method: str, response: Any) -> None:
        """Add a mock route with specified response.

        Args:
            path: URL path for the route
            method: HTTP method (GET, POST, etc.)
            response: Response data to return
        """
        raise NotImplementedError("To be implemented in Phase 2")

    def clear_routes(self) -> None:
        """Clear all configured routes."""
        raise NotImplementedError("To be implemented in Phase 2")
