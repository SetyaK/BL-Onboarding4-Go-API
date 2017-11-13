package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/SetyaK/BL-Onboarding3-Go-package"
	"github.com/SetyaK/BL-Onboarding3-Go-package/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"github.com/subosito/gotenv"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "ministore"
	app.Usage = "A CRUD command for ministore"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "Setya Kurniawan",
			Email: "setya.kurniawan@bukalapak.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "product",
			Usage: "doing CRUD with product",
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "get",
					Action: productGetAction,
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "id"},
					},
				},
				cli.Command{
					Name:   "list",
					Action: productListAction,
				},
				cli.Command{
					Name:   "add",
					Action: productAddAction,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "name"},
						cli.StringFlag{Name: "description"},
						cli.IntFlag{Name: "stock"},
					},
				},
				cli.Command{
					Name:   "update",
					Action: productUpdateAction,
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "id"},
						cli.StringFlag{Name: "name"},
						cli.StringFlag{Name: "description"},
					},
				},
				cli.Command{
					Name:   "delete",
					Action: productDeleteAction,
					Flags: []cli.Flag{
						cli.Int64Flag{Name: "id"},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func getProductRepository() ministore.ProductRepository {
	gotenv.Load()
	os.Setenv("DATABASE_ADAPTER", "mysql")

	// Initialize database session
	sess, err := database.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	return ministore.ProductRepository{Session: sess}
}

func productGetAction(c *cli.Context) error {
	pid := c.Int64("id")
	if pid < 1 {
		cli.ShowSubcommandHelp(c)
	} else {
		pr := getProductRepository()
		p, err := pr.GetByID(pid)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.App.Writer, " Product ID: %d\n Name: %s\n Description: %s\n Stock: %d", p.ProductID, p.Name, p.Description, p.Stock)
	}
	return nil
}

func productAddAction(c *cli.Context) error {
	pname := c.String("name")
	pdescription := c.String("description")
	pstock := c.Int("stock")
	if pname == "" || pdescription == "" {
		cli.ShowSubcommandHelp(c)
	} else {
		pr := getProductRepository()
		pid, err := pr.Add(pname, pdescription, pstock)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.App.Writer, "Product added with ID = %d", pid)
	}
	return nil
}

func productUpdateAction(c *cli.Context) error {
	pid := c.Int64("id")
	pname := c.String("name")
	pdescription := c.String("description")
	if pid < 1 || pname == "" || pdescription == "" {
		cli.ShowSubcommandHelp(c)
	} else {
		pr := getProductRepository()
		product, err := pr.GetByID(pid)
		if err != nil {
			return err
		}
		product.Name = pname
		product.Description = pdescription
		result, err := pr.Update(&product)
		if err != nil {
			return err
		}
		if result {
			fmt.Fprintf(c.App.Writer, "Product successfull updated")
		} else {
			fmt.Fprintf(c.App.Writer, "Product failed to update")
		}
	}
	return nil
}
func productDeleteAction(c *cli.Context) error {
	pid := c.Int64("id")
	if pid < 1 {
		cli.ShowSubcommandHelp(c)
	} else {
		pr := getProductRepository()
		result, err := pr.Delete(pid)
		if err != nil {
			return err
		}
		if result {
			fmt.Fprintf(c.App.Writer, "Product successfull deleted")
		} else {
			fmt.Fprintf(c.App.Writer, "Product failed to delete")
		}
	}
	return nil
}
func productListAction(c *cli.Context) error {
	pr := getProductRepository()
	products, _, err := pr.GetAll()
	if err != nil {
		return err
	}
	data := [][]string{}
	for _, p := range products {
		arr := []string{strconv.FormatInt(p.ProductID, 10), p.Name, p.Description, strconv.Itoa(p.Stock)}
		data = append(data, arr)
	}
	table := tablewriter.NewWriter(c.App.Writer)
	table.SetHeader([]string{"Product ID", "Name", "Description", "Stock"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return nil
}
