package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/suimon/internal/core/domain/host"
)

const (
	dashboardName  = "💧 SUIMON: PRESS Q or ESC TO QUIT"
	emptyRowHeight = 1
)

var (
	DashboardConfigDefault = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.BorderTitle(dashboardName),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.TitleColor(cell.ColorRed),
		container.TitleFocusedColor(cell.ColorRed),
		container.Focused(),
	}

	CellConfigDefault = []container.Option{
		container.Border(linestyle.Light),
		container.AlignVertical(align.VerticalMiddle),
		container.AlignHorizontal(align.HorizontalCenter),
		container.TitleColor(cell.ColorRed),
		container.FocusedColor(cell.ColorWhite),
	}
)

// GetColumnsConfig returns the columns configuration based on the specified dashboard type.
func GetColumnsConfig(dashboard enums.TableType) (ColumnsConfig, error) {
	switch dashboard {
	case enums.TableTypeNode:
		return ColumnsConfigNode, nil
	case enums.TableTypeValidator:
		return ColumnsConfigValidator, nil
	case enums.TableTypeRPC:
		return ColumnsConfigRPC, nil
	case enums.TableTypeGasPriceAndSubsidy:
		return ColumnsConfigSystemState, nil
	default:
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}
}

// GetColumnsValues returns the columns values based on the specified dashboard type and host.
func GetColumnsValues(dashboard enums.TableType, host domainhost.Host) (ColumnValues, error) {
	switch dashboard {
	case enums.TableTypeNode:
		return GetNodeColumnValues(host), nil
	case enums.TableTypeValidator:
		return GetValidatorColumnValues(host), nil
	case enums.TableTypeRPC:
		return GetRPCColumnValues(host), nil
	case enums.TableTypeGasPriceAndSubsidy:
		return GeSystemStateColumnValues(host)
	default:
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}
}

// GetRowsConfig returns the rows configuration based on the specified dashboard type.
func GetRowsConfig(dashboard enums.TableType) (RowsConfig, error) {
	switch dashboard {
	case enums.TableTypeNode:
		return RowsConfigNode, nil
	case enums.TableTypeValidator:
		return RowsConfigValidator, nil
	case enums.TableTypeRPC:
		return RowsConfigRPC, nil
	case enums.TableTypeGasPriceAndSubsidy:
		return RowsConfigSystemState, nil
	default:
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}
}

// GetCellsConfig returns the cells configuration based on the specified dashboard type.
func GetCellsConfig(dashboard enums.TableType) (CellsConfig, error) {
	switch dashboard {
	case enums.TableTypeNode:
		return CellsConfigNode, nil
	case enums.TableTypeValidator:
		return CellsConfigValidator, nil
	case enums.TableTypeRPC:
		return CellsConfigRPC, nil
	case enums.TableTypeGasPriceAndSubsidy:
		return CellsConfigSystemState, nil
	default:
		return nil, fmt.Errorf("unknown dashboard type: %v", dashboard)
	}
}

// GetColumns returns a slice of Columns based on the given ColumnsConfig and Cells.
func GetColumns(columnsConfig ColumnsConfig, cells Cells) (Columns, error) {
	columns := make(Columns, len(columnsConfig))

	for columnName, width := range columnsConfig {
		dashCell, ok := cells[columnName]
		if !ok {
			return nil, fmt.Errorf("failed to get cell for column %s", columnName)
		}

		columnPct := NewColumnPct(width, dashCell.GetWidget())

		columns[columnName] = columnPct
	}

	return columns, nil
}

// GetRows creates a new set of rows based on the configuration provided.
// It accepts a RowsConfig object that specifies the height and columns for each row,
// a Cells object that maps cell names to cell objects,
// and a Columns object that maps column names to column objects.
// It returns a Rows object and an error. The Rows object is a slice of Row objects,
// where each Row object represents a row in the terminal grid.
func GetRows(rowsConfig RowsConfig, cells Cells, columns Columns) (Rows, error) {
	rows := make(Rows, 0, len(rowsConfig))

	for _, rowConfig := range rowsConfig {
		rowColumns := make([]grid.Element, 0, len(rowConfig.Columns))

		for _, columnName := range rowConfig.Columns {
			column, ok := columns[columnName]
			if !ok {
				return nil, fmt.Errorf("failed to get column %s", columnName)
			}

			rowColumns = append(rowColumns, column)
		}

		rows = append(rows, NewRowPct(rowConfig.Height, rowColumns...))
	}

	// add empty row to limit last row height
	rows = append(rows, NewRowPct(emptyRowHeight))

	return rows, nil
}

// GetCells creates a new set of cells based on the configuration provided.
// It accepts a CellsConfig object that maps column names to cell names,
// and a terminal object that represents the terminal where the cells will be displayed.
// It returns a Cells object and an error. The Cells object is a map that maps
// column names to cell objects.
func GetCells(cellsConfig CellsConfig) (Cells, error) {
	cells := make(Cells, len(cellsConfig))

	for columnName, cellConfig := range cellsConfig {
		widget, err := newWidgetByColumnName(columnName, cellConfig.Color)
		if err != nil {
			return nil, err
		}

		dashCell, err := NewCell(cellConfig.Title, cellConfig.Color, widget)
		if err != nil {
			return nil, fmt.Errorf("failed to create new cell for %s: %w", columnName, err)
		}

		cells[columnName] = dashCell
	}

	return cells, nil
}
