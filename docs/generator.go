package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"docs/util"
	"github.com/tidwall/gjson"
)

func main() {
	var conf, doc, tmpl string
	flag.StringVar(&conf, "conf", "schema.json", "Schema file")
	flag.StringVar(&doc, "doc", "docs/dev/layer.md", "Markdown file")
	flag.StringVar(&tmpl, "tmpl", "docs/dev/doc.md", "Output file")
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(doc, os.O_APPEND, 0666)
	value := gjson.GetBytes(content, "properties")
	value.ForEach(func(key, value gjson.Result) bool {
		_, err = f.WriteString("\n# " + strings.Title(key.String()) + "\n" + "This section is of " + "{{.Properties." + strings.Title(key.String()) + ".Type}} type that tells about the basic information of the user and consists of the following sub-sections:\n") //print name of sections
		if err != nil {
			log.Fatal(err)
		}
		println(key.String())
		result := gjson.GetBytes(content, "properties."+key.String())
		result.ForEach(func(key1, value gjson.Result) bool {
			_, err = f.WriteString("##    " + key1.String() + "" + "\n") //print name of sub-sections
			if err != nil {
				log.Fatal(err)
			}
			print("  ")
			println(key1.String())

			if key1.String() == "properties" {
				result1 := gjson.GetBytes(content, "properties."+key.String()+"."+key1.String())
				result1.ForEach(func(key2, value gjson.Result) bool {
					_, err := f.WriteString("###        " + key2.String() + "" + "\nSub-section of type {{.Properties." + strings.Title(key.String()) + ".Properties." + strings.Title(key2.String()) + ".Type}}, used to specify the " + key2.String() + " of the person\nThe schema snippet can be shown below:\n\n       " + key2.String() + ":\n        " + value.String() + "\n\n") //print properties of subsections
					if err != nil {
						log.Fatal(err)
					}
					print("         ")
					println(key2.String())
					return true
				})
			}
			if key1.String() == "items" {
				result1 := gjson.GetBytes(content, "properties."+key.String()+".items.properties")
				result1.ForEach(func(key2, value gjson.Result) bool {
					_, err := f.WriteString("###        " + key2.String() + "" + "\nSub-section of type {{.Properties." + strings.Title(key.String()) + ".Items.Properties." + strings.Title(key2.String()) + ".Type}}, used to specify the " + key2.String() + " of the person\nThe schema snippet can be shown below:\n\n       " + key2.String() + ":\n        " + value.String() + "\n\n") //print items of subsections
					if err != nil {
						log.Fatal(err)
					}
					print("         ")
					println(key2.String())
					return true
				})
			}
			return true
		})
		return true
	})
	f.Close()
	var schema util.SchemaData
	if err := json.Unmarshal(content, &schema); err != nil {
		log.Fatal(err)
	}
	tpl, err := template.ParseFiles(doc)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err = tpl.Execute(file, schema); err != nil {
		log.Fatal(err)
	}
}
