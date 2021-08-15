package odoo

func (cmd *Command) GoToNextPage() bool {
	changed := cmd.GoToPage(cmd.Page + 1)
	cmd.updateOffset()
	return changed
}
func (cmd *Command) GoToPreviousPage() bool {
	changed := cmd.GoToPage(cmd.Page - 1)
	cmd.updateOffset()
	return changed
}
func (cmd *Command) GoToPage(page int) bool {
	changed := false
	if page > cmd.Pages {
		page = cmd.Pages
	}
	if page < 1 {
		page = 1
	}
	if cmd.Page != page {
		cmd.Page = page
		changed = true
	}
	cmd.updateOffset()
	return changed
}
func (cmd *Command) GoToFirstPage() bool {
	return cmd.GoToPage(1)
}
func (cmd *Command) GoToLastPage() bool {
	return cmd.GoToPage(cmd.Pages)
}
func (cmd *Command) updateOffset() {
	cmd.Offset = (cmd.Page - 1) * cmd.Limit
}
