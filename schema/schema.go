package schema

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Schema fields

type StandardField struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Format      string `json:"format,omitempty"`
}

type Location struct {
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
}

type Highlights struct {
	Type            string        `json:"type"`
	Description     string        `json:"description"`
	AdditionalItems bool          `json:"additionalItems"`
	Items           StandardField `json:"items"`
}

// Main properties
type Core struct {
	Type                 string `json:"type"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Properties           struct {
		Name              StandardField `json:"name"`
		Title             StandardField `json:"title"`
		Image             StandardField `json:"image"`
		Email             StandardField `json:"email"`
		Phone             StandardField `json:"phone"`
		URL               StandardField `json:"url"`
		Summary           StandardField `json:"summary"`
		CurrentLocation   Location      `json:"currentLocation"`
		PermanentLocation Location      `json:"permanentLocation"`
	} `json:"properties"`
}

type Work struct {
	Type            string `json:"type"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name        StandardField `json:"name"`
			Description StandardField `json:"description"`
			Position    StandardField `json:"position"`
			Location    Location      `json:"location"`
			URL         StandardField `json:"url"`
			StartDate   StandardField `json:"startDate"`
			EndDate     StandardField `json:"endDate"`
			Summary     StandardField `json:"summary"`
			Highlights  Highlights    `json:"highlights"`
		} `json:"properties"`
	} `json:"items"`
}

type Education struct {
	Type            string `json:"type"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Institution StandardField `json:"institution"`
			Location    Location      `json:"location"`
			Area        StandardField `json:"area"`
			StudyType   StandardField `json:"studyType"`
			StartDate   StandardField `json:"startDate"`
			EndDate     StandardField `json:"endDate"`
			Score       struct {
				Type                 string `json:"type"`
				AdditionalProperties bool   `json:"additionalProperties"`
				Properties           struct {
					Scoretype  StandardField `json:"scoretype"`
					Scorevalue StandardField `json:"scorevalue"`
				} `json:"properties"`
			} `json:"score"`
			Courses struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				Items           StandardField `json:"items"`
			} `json:"courses"`
			Honors struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				Items           StandardField `json:"items"`
			} `json:"honors"`
		} `json:"properties"`
	} `json:"items"`
}

type Volunteer struct {
	Type            string `json:"type"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Organization StandardField `json:"organization"`
			Position     StandardField `json:"position"`
			Location     Location      `json:"location"`
			URL          StandardField `json:"url"`
			StartDate    StandardField `json:"startDate"`
			EndDate      StandardField `json:"endDate"`
			Summary      StandardField `json:"summary"`
			Highlights   Highlights    `json:"highlights"`
		} `json:"properties"`
	} `json:"items"`
}

type Publications struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name        StandardField `json:"name"`
			Publisher   StandardField `json:"publisher"`
			ReleaseDate StandardField `json:"releaseDate"`
			Resources   struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				URL             StandardField `json:"url"`
				Label           StandardField `json:"label"`
			} `json:"resources"`
			URL     StandardField `json:"url"`
			Summary StandardField `json:"summary"`
		} `json:"properties"`
	} `json:"items"`
}

type Legal struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name            StandardField `json:"name"`
			LegalType       StandardField `json:"legalType"`
			Description     StandardField `json:"description"`
			ApplicationDate StandardField `json:"applicationDate"`
			GrantDate       StandardField `json:"grantDate"`
			EndDate         StandardField `json:"endDate"`
			Resources       struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				URL             StandardField `json:"url"`
				Label           StandardField `json:"label"`
			} `json:"resources"`
			IDNumber StandardField `json:"idNumber"`
		} `json:"properties"`
	} `json:"items"`
}

type Skills struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name    StandardField `json:"name"`
			Keyword struct {
				Type            string `json:"type"`
				Description     string `json:"description"`
				AdditionalItems bool   `json:"additionalItems"`
				Items           struct {
					Type                 string `json:"type"`
					AdditionalProperties bool   `json:"additionalProperties"`
					Properties           struct {
						Name  StandardField `json:"name"`
						Score StandardField `json:"score"`
					} `json:"properties"`
				} `json:"items"`
			} `json:"keyword"`
		} `json:"properties"`
	} `json:"items"`
}

type Awards struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Title   StandardField `json:"title"`
			Date    StandardField `json:"date"`
			Awarder StandardField `json:"awarder"`
			Summary StandardField `json:"summary"`
		} `json:"properties"`
	} `json:"items"`
}

type Projects struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name        StandardField `json:"name"`
			Location    Location      `json:"location"`
			Description StandardField `json:"description"`
			Highlights  Highlights    `json:"highlights"`
			Keywords    struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				Items           StandardField `json:"items"`
			} `json:"keywords"`
			StartDate StandardField `json:"startDate"`
			EndDate   StandardField `json:"endDate"`
			Resources struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				URL             StandardField `json:"url"`
				Label           StandardField `json:"label"`
			} `json:"resources"`
			URL   StandardField `json:"url"`
			Roles struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				Items           StandardField `json:"items"`
			} `json:"roles"`
			Entity StandardField `json:"entity"`
			Type   StandardField `json:"type"`
		} `json:"properties"`
	} `json:"items"`
}

type Certificates struct {
	Type            string `json:"type"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Code          StandardField `json:"code"`
			Name          StandardField `json:"name"`
			Website       StandardField `json:"website"`
			Verification  StandardField `json:"verification"`
			GrantDate     StandardField `json:"grantDate"`
			Score         StandardField `json:"score"`
			EndDate       StandardField `json:"endDate"`
			DoesNotExpire struct {
				Type   string `json:"type"`
				Format string `json:"format"`
			} `json:"doesNotExpire"`
		} `json:"properties"`
	} `json:"items"`
}

type References struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name      StandardField `json:"name"`
			Company   StandardField `json:"company"`
			Position  StandardField `json:"position"`
			Reference StandardField `json:"reference"`
		} `json:"properties"`
	} `json:"items"`
}

type Languages struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Language StandardField `json:"language"`
			Score    StandardField `json:"score"`
		} `json:"properties"`
	} `json:"items"`
}

type Interests struct {
	Type            string `json:"type"`
	AdditionalItems bool   `json:"additionalItems"`
	Items           struct {
		Type                 string `json:"type"`
		AdditionalProperties bool   `json:"additionalProperties"`
		Properties           struct {
			Name     StandardField `json:"name"`
			Keywords struct {
				Type            string        `json:"type"`
				Description     string        `json:"description"`
				AdditionalItems bool          `json:"additionalItems"`
				Items           StandardField `json:"items"`
			} `json:"keywords"`
		} `json:"properties"`
	} `json:"items"`
}

type Meta struct {
	Type                 string `json:"type"`
	Description          string `json:"description"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Properties           struct {
		Canonical    StandardField `json:"canonical"`
		Version      StandardField `json:"version"`
		LastModified StandardField `json:"lastModified"`
	} `json:"properties"`
}

type Schema struct {
	Schema               string `json:"$schema"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Properties           struct {
		Core         Core         `json:"core"`
		Work         Work         `json:"work"`
		Education    Education    `json:"education"`
		Volunteer    Volunteer    `json:"volunteer"`
		Publications Publications `json:"publications"`
		Legal        Legal        `json:"legal"`
		Skills       Skills       `json:"skills"`
		Awards       Awards       `json:"awards"`
		Projects     Projects     `json:"projects"`
		Certificates Certificates `json:"certificates"`
		References   References   `json:"references"`
		Languages    Languages    `json:"languages"`
		Interests    Interests    `json:"interests"`
		Meta         Meta         `json:"meta"`
	} `json:"properties"`
}

// GetSchema fills the Schema{} struct with correct values and returns it
func GetSchema() (Schema, error) {
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
