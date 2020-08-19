package util

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func IsWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}

func IsAssistant(obj glib.IObject) (*gtk.Assistant, error) {
	if win, ok := obj.(*gtk.Assistant); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Assistant")
}

func IsDialog(obj glib.IObject) (*gtk.Dialog, error) {
	if dialog, ok := obj.(*gtk.Dialog); ok {
		return dialog, nil
	}
	return nil, errors.New("not a *gtk.Dialog")
}

func IsEntry(obj glib.IObject) (*gtk.Entry, error) {
	if entry, ok := obj.(*gtk.Entry); ok {
		return entry, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func IsComboBoxText(obj glib.IObject) (*gtk.ComboBoxText, error) {
	if comboBoxText, ok := obj.(*gtk.ComboBoxText); ok {
		return comboBoxText, nil
	}
	return nil, errors.New("not a *gtk.ComboBoxText")
}

func IsLabel(obj glib.IObject) (*gtk.Label, error) {
	if label, ok := obj.(*gtk.Label); ok {
		return label, nil
	}
	return nil, errors.New("not a *gtk.Label")
}

func IsSpinButton(obj glib.IObject) (*gtk.SpinButton, error) {
	if spinBtn, ok := obj.(*gtk.SpinButton); ok {
		return spinBtn, nil
	}
	return nil, errors.New("not a *gtk.SpinButton")
}

func IsPopover(obj glib.IObject) (*gtk.Popover, error) {
	if popover, ok := obj.(*gtk.Popover); ok {
		return popover, nil
	}
	return nil, errors.New("not a *gtk.Popover")
}

func IsButton(obj glib.IObject) (*gtk.Button, error) {
	if btn, ok := obj.(*gtk.Button); ok {
		return btn, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func IsSearchEntry(obj glib.IObject) (*gtk.SearchEntry, error) {
	if searchEntry, ok := obj.(*gtk.SearchEntry); ok {
		return searchEntry, nil
	}
	return nil, errors.New("not a *gtk.SearchEntry")
}

func IsListStore(obj glib.IObject) (*gtk.ListStore, error) {
	if listStore, ok := obj.(*gtk.ListStore); ok {
		return listStore, nil
	}
	return nil, errors.New("not a *gtk.ListStore")
}

func IsEntryCompletion(obj glib.IObject) (*gtk.EntryCompletion, error) {
	if entryCompletion, ok := obj.(*gtk.EntryCompletion); ok {
		return entryCompletion, nil
	}
	return nil, errors.New("not a *gtk.EntryCompletion")
}

func IsSpinner(obj glib.IObject) (*gtk.Spinner, error) {
	if spinner, ok := obj.(*gtk.Spinner); ok {
		return spinner, nil
	}
	return nil, errors.New("not a *gtk.Spinner")
}

func IsImage(obj glib.IObject) (*gtk.Image, error) {
	if image, ok := obj.(*gtk.Image); ok {
		return image, nil
	}
	return nil, errors.New("not a *gtk.Image")
}

func IsGrid(obj glib.IObject) (*gtk.Grid, error) {
	if grid, ok := obj.(*gtk.Grid); ok {
		return grid, nil
	}
	return nil, errors.New("not a *gtk.Grid")
}

func IsTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	if treeView, ok := obj.(*gtk.TreeView); ok {
		return treeView, nil
	}
	return nil, errors.New("not a *gtk.TreeView")
}

func IsBox(obj glib.IObject) (*gtk.Box, error) {
	if box, ok := obj.(*gtk.Box); ok {
		return box, nil
	}
	return nil, errors.New("not a *gtk.Box")
}

func IsPopoverMenu(obj glib.IObject) (*gtk.PopoverMenu, error) {
	if popoverMenu, ok := obj.(*gtk.PopoverMenu); ok {
		return popoverMenu, nil
	}
	return nil, errors.New("not a *gtk.PopoverMenu")
}

func IsCellRenderer(obj glib.IObject) (*gtk.CellRenderer, error) {
	if cellRenderer, ok := obj.(*gtk.CellRenderer); ok {
		return cellRenderer, nil
	}
	return nil, errors.New("not a *gtk.CellRenderer")
}

func IsCellRendererText(obj glib.IObject) (*gtk.CellRendererText, error) {
	if cellRendererText, ok := obj.(*gtk.CellRendererText); ok {
		return cellRendererText, nil
	}
	return nil, errors.New("not a *gtk.CellRendererText")
}

func IsDrawingArea(obj glib.IObject) (*gtk.DrawingArea, error) {
	if drawingArea, ok := obj.(*gtk.DrawingArea); ok {
		return drawingArea, nil
	}
	return nil, errors.New("not a *gtk.DrawingArea")
}
