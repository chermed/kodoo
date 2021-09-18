package odoo

func (cmd *Command) GoToNextPage(rotate bool) bool {
	changed := cmd.GoToPage(cmd.Page+1, rotate)
	cmd.updateOffset()
	return changed
}
func (cmd *Command) GoToPreviousPage(rotate bool) bool {
	changed := cmd.GoToPage(cmd.Page-1, rotate)
	cmd.updateOffset()
	return changed
}
func (cmd *Command) GoToPage(page int, rotate bool) bool {
	changed := false
	if rotate {
		changed = true
	}
	if page > cmd.Pages {
		if rotate {
			page = 1
		} else {
			page = cmd.Pages
		}
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
func (cmd *Command) GoToFirstPage(rotate bool) bool {
	return cmd.GoToPage(1, rotate)
}
func (cmd *Command) GoToLastPage(rotate bool) bool {
	return cmd.GoToPage(cmd.Pages, rotate)
}
func (cmd *Command) updateOffset() {
	cmd.Offset = (cmd.Page - 1) * cmd.Limit
}
