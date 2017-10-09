package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/tidwall/gjson"
)

func main() {
	var conf, doc, tmpl string
	flag.StringVar(&conf, "conf", "schema.json", "Schema file")
	flag.StringVar(&doc, "doc", "docs/doc.md", "Markdown file")
	flag.StringVar(&tmpl, "tmpl", "docs/user/test.md", "Output file")
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(doc, os.O_APPEND, 0666)
	value := gjson.GetBytes(content, "properties")
	value.ForEach(func(key, value gjson.Result) bool {
		_, err = f.WriteString("\n# " + strings.Title(key.String()) + "\n" + "This section is of " + "{{.Properties." + strings.Title(key.String()) + ".Title.Description}} type that tells about the basic information of the user and consists of the following sub-sections:\n") //print name of sections
		if err != nil {
			log.Fatal(err)
		}
		println(key.String())
		result := gjson.GetBytes(content, "properties."+key.String())
		result.ForEach(func(key1, value gjson.Result) bool {
			_, err = f.WriteString("##    {{" + key1.String() + "}}" + "\n") //print name of sub-sections
			if err != nil {
				log.Fatal(err)
			}
			print("  ")
			println(key1.String())

			if key1.String() == "properties" {
				result1 := gjson.GetBytes(content, "properties."+key.String()+"."+key1.String())
				result1.ForEach(func(key2, value gjson.Result) bool {
					_, err := f.WriteString("###        {{" + key2.String() + "}}" + "\nSub-section of type {{.Properties." + strings.Title(key.String()) + ".Properties." + strings.Title(key2.String()) + ".Type}},used to specify the " + key2.String() + " of the person\nThe schema snippet can be shown below:\n        " + key2.String() + ":\n        " + value.String() + "\n") //print properties of subsections
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
					_, err := f.WriteString("###        {{" + key2.String() + "}}" + "\nSub-section of type {{.Properties." + strings.Title(key.String()) + ".Items.Properties." + strings.Title(key2.String()) + ".Type}},used to specify the " + key2.String() + " of the person\nThe schema snippet can be shown below:\n        " + key2.String() + ":\n        " + value.String() + "\n") //print items of subsections
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
	var schema SchemaData
	if err := json.Unmarshal(content, &schema); err != nil {
		log.Fatal(err)
	}
	fmt.Println(schema.Properties.Core.Type) //check
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

type SchemaData struct {
	Schema               string `json:"$schema"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Properties           struct {
		Core struct {
			Type                 string `json:"type"`
			AdditionalProperties bool   `json:"additionalProperties"`
			Properties           struct {
				Name struct {
					Type string `json:"type"`
				} `json:"name"`
				Title struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"title"`
				Image struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"image"`
				Email struct {
					Type        string `json:"type"`
					Description string `json:"description"`
					Format      string `json:"format"`
				} `json:"email"`
				Phone struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"phone"`
				URL struct {
					Type        string `json:"type"`
					Format      string `json:"format"`
					Description string `json:"description"`
				} `json:"url"`
				Summary struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"summary"`
				CurrentLocation struct {
					Type        string `json:"type"`
					Format      string `json:"format"`
					Description string `json:"description"`
					Properties  struct {
						Lat struct {
							Type string `json:"type"`
						} `json:"lat"`
						Long struct {
							Type string `json:"type"`
						} `json:"long"`
					} `json:"properties"`
				} `json:"currentLocation"`
				PermanentLocation struct {
					Type        string `json:"type"`
					Format      string `json:"format"`
					Description string `json:"description"`
					Properties  struct {
						Lat struct {
							Type string `json:"type"`
						} `json:"lat"`
						Long struct {
							Type string `json:"type"`
						} `json:"long"`
					} `json:"properties"`
				} `json:"permanentLocation"`
			} `json:"properties"`
		} `json:"core"`
		Work struct {
			Type            string `json:"type"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Description struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"description"`
					Position struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"position"`
					Location struct {
						Type        string `json:"type"`
						Format      string `json:"format"`
						Description string `json:"description"`
						Properties  struct {
							Lat struct {
								Type string `json:"type"`
							} `json:"lat"`
							Long struct {
								Type string `json:"type"`
							} `json:"long"`
						} `json:"properties"`
					} `json:"location"`
					URL struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"url"`
					StartDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"startDate"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					Summary struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"summary"`
					Highlights struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"highlights"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"work"`
		Education struct {
			Type            string `json:"type"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Institution struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"institution"`
					Area struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"area"`
					StudyType struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"studyType"`
					StartDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"startDate"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					Score struct {
						Type                 string `json:"type"`
						AdditionalProperties bool   `json:"additionalProperties"`
						Properties           struct {
							Scoretype struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"scoretype"`
							Scorevalue struct {
								Type        string `json:"type"`
								Description string `json:"description"`
							} `json:"scorevalue"`
						} `json:"properties"`
					} `json:"score"`
					Courses struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"courses"`
					Honors struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"honors"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"education"`
		Volunteer struct {
			Type            string `json:"type"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Organization struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"organization"`
					Position struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"position"`
					Location struct {
						Type        string `json:"type"`
						Format      string `json:"format"`
						Description string `json:"description"`
						Properties  struct {
							Lat struct {
								Type string `json:"type"`
							} `json:"lat"`
							Long struct {
								Type string `json:"type"`
							} `json:"long"`
						} `json:"properties"`
					} `json:"location"`
					URL struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"url"`
					StartDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"startDate"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					Summary struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"summary"`
					Highlights struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"highlights"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"volunteer"`
		Publications struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Publisher struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"publisher"`
					ReleaseDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"releaseDate"`
					Resources struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						URL             struct {
							Type        string `json:"type"`
							Format      string `json:"format"`
							Description string `json:"description"`
						} `json:"url"`
						Label struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"label"`
					} `json:"resources"`
					URL struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"url"`
					Summary struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"summary"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"publications"`
		Legal struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					LegalType struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"legalType"`
					Description struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"description"`
					ApplicationDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"applicationDate"`
					GrantDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"grantDate"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					Resources struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						URL             struct {
							Type        string `json:"type"`
							Format      string `json:"format"`
							Description string `json:"description"`
						} `json:"url"`
						Label struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"label"`
					} `json:"resources"`
					IDNumber struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"idNumber"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"legal"`
		Skills struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Keyword struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type                 string `json:"type"`
							AdditionalProperties bool   `json:"additionalProperties"`
							Properties           struct {
								Name struct {
									Type        string `json:"type"`
									Description string `json:"description"`
								} `json:"name"`
								Score struct {
									Type        string `json:"type"`
									Description string `json:"description"`
								} `json:"score"`
							} `json:"properties"`
						} `json:"items"`
					} `json:"keyword"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"skills"`
		Awards struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Title struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"title"`
					Date struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"date"`
					Awarder struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"awarder"`
					Summary struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"summary"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"awards"`
		Projects struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Location struct {
						Type        string `json:"type"`
						Format      string `json:"format"`
						Description string `json:"description"`
						Properties  struct {
							Lat struct {
								Type string `json:"type"`
							} `json:"lat"`
							Long struct {
								Type string `json:"type"`
							} `json:"long"`
						} `json:"properties"`
					} `json:"location"`
					Description struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"description"`
					Highlights struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"highlights"`
					Keywords struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"keywords"`
					StartDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"startDate"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					Resources struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						URL             struct {
							Type        string `json:"type"`
							Format      string `json:"format"`
							Description string `json:"description"`
						} `json:"url"`
						Label struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"label"`
					} `json:"resources"`
					URL struct {
						Type        string `json:"type"`
						Format      string `json:"format"`
						Description string `json:"description"`
					} `json:"url"`
					Roles struct {
						Type            string `json:"type"`
						Description     string `json:"description"`
						AdditionalItems bool   `json:"additionalItems"`
						Items           struct {
							Type        string `json:"type"`
							Description string `json:"description"`
						} `json:"items"`
					} `json:"roles"`
					Entity struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"entity"`
					Type struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"type"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"projects"`
		Certificates struct {
			Type            string `json:"type"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Code struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"code"`
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Website struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"website"`
					Verification struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"verification"`
					GrantDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"grantDate"`
					Score struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"score"`
					EndDate struct {
						Type        string `json:"type"`
						Description string `json:"description"`
						Format      string `json:"format"`
					} `json:"endDate"`
					DoesNotExpire struct {
						Type   string `json:"type"`
						Format string `json:"format"`
					} `json:"doesNotExpire"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"certificates"`
		References struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
					Company struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"company"`
					Position struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"position"`
					Reference struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"reference"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"references"`
		Languages struct {
			Type            string `json:"type"`
			Description     string `json:"description"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Language struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"language"`
					Score struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"score"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"languages"`
		Interests struct {
			Type            string `json:"type"`
			AdditionalItems bool   `json:"additionalItems"`
			Items           struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Name struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"name"`
				} `json:"properties"`
			} `json:"items"`
		} `json:"interests"`
		Meta struct {
			Type                 string `json:"type"`
			Description          string `json:"description"`
			AdditionalProperties bool   `json:"additionalProperties"`
			Properties           struct {
				Canonical struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"canonical"`
				Version struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"version"`
				LastModified struct {
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"lastModified"`
			} `json:"properties"`
		} `json:"meta"`
	} `json:"properties"`
}
