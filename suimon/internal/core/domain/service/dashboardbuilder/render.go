package dashboardbuilder

import (
	"fmt"
	"os"
	"time"

	"github.com/mum4k/termdash"
	termcell "github.com/mum4k/termdash/cell"
	"golang.org/x/sync/errgroup"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
)

const (
	renderInterval = 200 * time.Millisecond
	queryInterval  = 1 * time.Second
)

// Render displays the dashboard on the terminal and updates the cells with new data periodically.
func (db *Builder) Render() (err error) {
	// Use a deferred function to call db.TearDown() if there were errors or panics
	defer func() {
		if err != nil {
			db.tearDown()
		}

		if err := recover(); err != nil {
			// Handle the panic by logging the error and exiting the program
			db.tearDown()

			db.cliGateway.Error(fmt.Sprintf("panic: %v", err))

			os.Exit(1)
		}
	}()

	var errGroup errgroup.Group

	tickerQuery := time.NewTicker(queryInterval)
	defer tickerQuery.Stop()

	// Start a goroutine for the metric retrieval loop
	errGroup.Go(func() error {
		for {
			select {
			case <-tickerQuery.C:
				if err := db.host.GetMetrics(); err != nil {
					return err
				}
			case <-db.ctx.Done():
				return nil
			}
		}
	})

	tickerRerender := time.NewTicker(renderInterval)
	defer tickerRerender.Stop()

	// Start a goroutine for the dashboard rendering loop
	errGroup.Go(func() error {
		for {
			select {
			case <-tickerRerender.C:
				columns := dashboards.GetNodeColumnValues(db.host)

				for columnName, cell := range db.cells {
					columnValue, ok := columns[columnName]
					if !ok {
						return fmt.Errorf("failed to get metric for column %s", columnName)
					}

					if err := cell.Write(columnValue, []termcell.Option{termcell.FgColor(termcell.ColorDefault)}); err != nil {
						return err
					}
				}
			case <-db.ctx.Done():
				return nil
			}
		}
	})

	// Display the dashboard on the terminal and handle errors
	if err := termdash.Run(
		db.ctx, db.terminal, db.dashboard,
		termdash.KeyboardSubscriber(db.quitter),
	); err != nil {
		return fmt.Errorf("failed to run terminal dashboard: %w", err)
	}

	return errGroup.Wait()
}
