package daemon

import (
	"fmt"
	"strings"
)

// BuildPrompt constructs the task prompt for an agent CLI.
// The prompt is intentionally minimal — it provides only the issue ID and
// instructs the agent to use the multica CLI to fetch details on demand.
// Skill instructions are injected via the runtime's native config mechanism
// (e.g., .claude/CLAUDE.md, AGENTS.md) by execenv.InjectRuntimeConfig.
func BuildPrompt(task Task) string {
	var b strings.Builder
	b.WriteString("You are running as a local coding agent for a Multica workspace.\n\n")

	fmt.Fprintf(&b, "Your assigned issue ID is: %s\n\n", task.IssueID)

	b.WriteString("Use the `multica` CLI to fetch the issue details and any context you need:\n\n")
	fmt.Fprintf(&b, "  multica issue get %s --output json    # Full issue details\n", task.IssueID)
	fmt.Fprintf(&b, "  multica issue comment list %s         # Comments and discussion\n\n", task.IssueID)

	b.WriteString("Workflow:\n")
	fmt.Fprintf(&b, "1. Run `multica issue get %s --output json` to understand your task\n", task.IssueID)
	fmt.Fprintf(&b, "2. Run `multica issue status %s in_progress` to mark it as started\n", task.IssueID)
	b.WriteString("3. Complete the work in the local codebase\n")
	fmt.Fprintf(&b, "4. Run `multica issue status %s done` (or `in_review` if human review is needed)\n", task.IssueID)
	fmt.Fprintf(&b, "5. If blocked, run `multica issue status %s blocked` and post a comment explaining why\n\n", task.IssueID)

	return b.String()
}
