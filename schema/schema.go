package schema

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Schema struct {
	Schema               string
	Title                string
	Type                 string
	AdditionalProperties bool
	Properties           struct {
		Core struct {
			Type                 string
			AdditionalProperties bool
			Properties           struct {
				Name struct {
					Type        string
					Description string
				}
				Title struct {
					Type        string
					Description string
				}
				Image struct {
					Type        string
					Description string
				}
				Email struct {
					Type        string
					Description string
					Format      string
				}
				Phone struct {
					Type        string
					Description string
				}
				URL struct {
					Type        string
					Format      string
					Description string
				}
				Summary struct {
					Type        string
					Format      string
					Description string
				}
				CurrentLocation struct {
					Type        string
					Format      string
					Description string
					Properties  struct {
						Lat struct {
							Type string
						}
						Long struct {
							Type string
						}
					}
				}
				PermanentLocation struct {
					Type        string
					Format      string
					Description string
					Properties  struct {
						Lat struct {
							Type string
						}
						Long struct {
							Type string
						}
					}
				}
			}
		}
		Work struct {
			Type            string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					Description struct {
						Type        string
						Description string
					}
					Position struct {
						Type        string
						Description string
					}
					Location struct {
						Type        string
						Format      string
						Description string
						Properties  struct {
							Lat struct {
								Type string
							}
							Long struct {
								Type string
							}
						}
					}
					URL struct {
						Type        string
						Description string
						Format      string
					}
					StartDate struct {
						Type        string
						Description string
						Format      string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					Summary struct {
						Type        string
						Description string
					}
					Highlights struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
				}
			}
		}
		Education struct {
			Type            string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Institution struct {
						Type        string
						Description string
					}
					Location struct {
						Type        string
						Format      string
						Description string
						Properties  struct {
							Lat struct {
								Type string
							}
							Long struct {
								Type string
							}
						}
					}
					Area struct {
						Type        string
						Description string
					}
					StudyType struct {
						Type        string
						Description string
					}
					StartDate struct {
						Type        string
						Description string
						Format      string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					Score struct {
						Type                 string
						AdditionalProperties bool
						Properties           struct {
							Scoretype struct {
								Type        string
								Description string
							}
							Scorevalue struct {
								Type        string
								Description string
							}
						}
					}
					Courses struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
					Honors struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
				}
			}
		}
		Volunteer struct {
			Type            string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Organization struct {
						Type        string
						Description string
					}
					Position struct {
						Type        string
						Description string
					}
					Location struct {
						Type        string
						Format      string
						Description string
						Properties  struct {
							Lat struct {
								Type string
							}
							Long struct {
								Type string
							}
						}
					}
					URL struct {
						Type        string
						Description string
						Format      string
					}
					StartDate struct {
						Type        string
						Description string
						Format      string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					Summary struct {
						Type        string
						Description string
					}
					Highlights struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
				}
			}
		}
		Publications struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					Publisher struct {
						Type        string
						Description string
					}
					ReleaseDate struct {
						Type        string
						Description string
					}
					Resources struct {
						Type            string
						Description     string
						AdditionalItems bool
						URL             struct {
							Type        string
							Format      string
							Description string
						}
						Label struct {
							Type        string
							Description string
						}
					}
					URL struct {
						Type        string
						Format      string
						Description string
					}
					Summary struct {
						Type        string
						Description string
					}
				}
			}
		}
		Legal struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					LegalType struct {
						Type        string
						Description string
					}
					Description struct {
						Type        string
						Description string
					}
					ApplicationDate struct {
						Type        string
						Description string
						Format      string
					}
					GrantDate struct {
						Type        string
						Description string
						Format      string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					Resources struct {
						Type            string
						Description     string
						AdditionalItems bool
						URL             struct {
							Type        string
							Format      string
							Description string
						}
						Label struct {
							Type        string
							Description string
						}
					}
					IDNumber struct {
						Type        string
						Description string
					}
				}
			}
		}
		Skills struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					Keyword struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type                 string
							AdditionalProperties bool
							Properties           struct {
								Name struct {
									Type        string
									Description string
								}
								Score struct {
									Type        string
									Description string
								}
							}
						}
					}
				}
			}
		}
		Awards struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Title struct {
						Type        string
						Description string
					}
					Date struct {
						Type        string
						Description string
						Format      string
					}
					Awarder struct {
						Type        string
						Description string
					}
					Summary struct {
						Type        string
						Description string
					}
				}
			}
		}
		Projects struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					Location struct {
						Type        string
						Format      string
						Description string
						Properties  struct {
							Lat struct {
								Type string
							}
							Long struct {
								Type string
							}
						}
					}
					Description struct {
						Type        string
						Description string
					}
					Highlights struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
					Keywords struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
					StartDate struct {
						Type        string
						Description string
						Format      string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					Resources struct {
						Type            string
						Description     string
						AdditionalItems bool
						URL             struct {
							Type        string
							Format      string
							Description string
						}
						Label struct {
							Type        string
							Description string
						}
					}
					URL struct {
						Type        string
						Format      string
						Description string
					}
					Roles struct {
						Type            string
						Description     string
						AdditionalItems bool
						Items           struct {
							Type        string
							Description string
						}
					}
					Entity struct {
						Type        string
						Description string
					}
					Type struct {
						Type        string
						Description string
					}
				}
			}
		}
		Certificates struct {
			Type            string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Code struct {
						Type        string
						Description string
					}
					Name struct {
						Type        string
						Description string
					}
					Website struct {
						Type        string
						Description string
						Format      string
					}
					Verification struct {
						Type        string
						Description string
						Format      string
					}
					GrantDate struct {
						Type        string
						Description string
						Format      string
					}
					Score struct {
						Type        string
						Description string
					}
					EndDate struct {
						Type        string
						Description string
						Format      string
					}
					DoesNotExpire struct {
						Type   string
						Format string
					}
				}
			}
		}
		References struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
					Company struct {
						Type        string
						Description string
					}
					Position struct {
						Type        string
						Description string
					}
					Reference struct {
						Type        string
						Description string
					}
				}
			}
		}
		Languages struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Language struct {
						Type        string
						Description string
					}
					Score struct {
						Type        string
						Description string
					}
				}
			}
		}
		Interests struct {
			Type            string
			Description     string
			AdditionalItems bool
			Items           struct {
				Type                 string
				AdditionalProperties bool
				Properties           struct {
					Name struct {
						Type        string
						Description string
					}
				}
				Keywords struct {
					Type            string
					Description     string
					AdditionalItems bool
					Items           struct {
						Type        string
						Description string
					}
				}
			}
		}
		Meta struct {
			Type                 string
			Description          string
			AdditionalProperties bool
			Properties           struct {
				Canonical struct {
					Type        string
					Description string
				}
				Version struct {
					Type        string
					Description string
				}
				LastModified struct {
					Type        string
					Description string
					Format      string
				}
			}
		}
	}
}

func GetDefaultSchema() (Schema, error) {
	valuesFile, err := os.Open("./schema.json")
	if err != nil {
		return Schema{}, err
	}
	defer valuesFile.Close()

	values, err := ioutil.ReadAll(valuesFile)
	if err != nil {
		return Schema{}, err
	}
	schema := Schema{}
	err = json.Unmarshal(values, &schema)
	return schema, nil
}

func GetSchema() Schema {
	return Schema{}
}
