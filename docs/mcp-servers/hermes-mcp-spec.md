# Hermes MCP Server Specification

**Purpose:** Provide AI agents with access to work state operations.

## Overview

The Hermes MCP server enables agents to:
1. Create and manage beads issues
2. Query work state (ready, blocked, in_progress)
3. Sync work with git
4. Track dependencies and blockers

**Foundational:** Hermes is part of the core trinity (Athena, Hephaestus, Hermes).

## Tools

### create_issue

Create a new beads issue.

```json
{
  "name": "create_issue",
  "description": "Create a new beads issue",
  "parameters": {
    "title": {
      "type": "string",
      "description": "Issue title"
    },
    "type": {
      "type": "string",
      "enum": ["task", "feature", "bug", "epic"],
      "default": "task"
    },
    "priority": {
      "type": "string",
      "enum": ["P1", "P2", "P3"],
      "default": "P2"
    },
    "description": {
      "type": "string",
      "description": "Detailed description"
    },
    "parent": {
      "type": "string",
      "description": "Parent issue ID (for epics)"
    },
    "depends_on": {
      "type": "array",
      "items": {"type": "string"},
      "description": "Issues this depends on"
    }
  }
}
```

### update_issue

Update an existing issue.

```json
{
  "name": "update_issue",
  "description": "Update a beads issue",
  "parameters": {
    "id": {
      "type": "string",
      "description": "Issue ID"
    },
    "status": {
      "type": "string",
      "enum": ["open", "in_progress", "closed"]
    },
    "priority": {
      "type": "string",
      "enum": ["P1", "P2", "P3"]
    },
    "notes": {
      "type": "string",
      "description": "Additional notes to append"
    }
  }
}
```

### get_ready

Find issues ready for work (no blockers).

```json
{
  "name": "get_ready",
  "description": "Get issues with no blockers",
  "parameters": {
    "parent": {
      "type": "string",
      "description": "Filter by parent epic ID"
    },
    "limit": {
      "type": "integer",
      "default": 10
    },
    "priority": {
      "type": "string",
      "enum": ["P1", "P2", "P3"],
      "description": "Filter by minimum priority"
    }
  }
}
```

### get_blocked

Find blocked issues and their blockers.

```json
{
  "name": "get_blocked",
  "description": "Get blocked issues with their dependencies",
  "parameters": {
    "parent": {
      "type": "string",
      "description": "Filter by parent epic ID"
    }
  }
}
```

### show_issue

Get full details for an issue.

```json
{
  "name": "show_issue",
  "description": "Get complete issue details",
  "parameters": {
    "id": {
      "type": "string",
      "description": "Issue ID"
    }
  }
}
```

### close_issue

Close an issue with reason.

```json
{
  "name": "close_issue",
  "description": "Close a beads issue",
  "parameters": {
    "id": {
      "type": "string",
      "description": "Issue ID"
    },
    "reason": {
      "type": "string",
      "description": "Closure reason/summary"
    }
  }
}
```

### sync

Sync beads state with git.

```json
{
  "name": "sync",
  "description": "Sync beads changes to git",
  "parameters": {}
}
```

### get_molecule_progress

Get progress for a molecule (poured formula).

```json
{
  "name": "get_molecule_progress",
  "description": "Get progress summary for a molecule",
  "parameters": {
    "id": {
      "type": "string",
      "description": "Molecule root issue ID"
    }
  }
}
```

## Integration Notes

### With Hephaestus

When Hephaestus dispatches work:
1. Call `update_issue` to set status=in_progress
2. On completion, call `close_issue`
3. Call `sync` after state changes

### With Hephaestus (Automaton Completion)

> **Note:** Cyclopes was consolidated into Hephaestus (2026-01-19).

When Automaton completes:
1. Call `close_issue` for the work item
2. Log completion to Chronicle

## Implementation Location

The MCP server should be added to:
- Repository: `beads`
- Path: `integrations/beads-mcp/`
- Existing: Python MCP server already exists

## See Also

- [Beads CLI Reference](https://github.com/steveyegge/beads)
- [Molecule/Formula Documentation](../../.agents/formulas/)
