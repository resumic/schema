package schema

type locationSchema struct {
	Lat  float64 `json:"lat" schema:"example:35.712758"`
	Long float64 `json:"long" schema:"example:51.392114"`
}

type scoreSchema struct {
	Type  string `json:"type" schema:"description:type of the score;example:GPA"`
	Value string `json:"value" schema:"description:value of the score;example:3.4"`
	Best  string `json:"best" schema:"description:best possible score;example:4"`
	Worst string `json:"worst" schema:"description:worst possible score;example:0"`
}

type resourceSchema struct {
	URL   string `json:"url" schema:"description:url of the resource;format:uri;example:http://www.example.com/my-example-slides/"`
	Label string `json:"label" schema:"description:label of the resource;example:Slides;example:Slides"`
}

type profileSchema struct {
	Network  string `json:"network" schema:"description:name of the network;example:github"`
	Username string `json:"username" schema:"description:username in the network;example:john_doe"`
	URL      string `json:"url" schema:"description:url of the profile;format:uri;example:https://github.com/john_doe"`
}

type coreSchema struct {
	Title      string `json:"title,omitempty" schema:"description:job title;example:Software Engineer"`
	WorkArea   string `json:"workArea,omitempty" schema:"description:work area which could be city or country or even remote;example:Munich Citycenter + 20kms"`
	LivingArea string `json:"livingArea,omitempty" schema:"description:living area which could be city or country or even continent;example:Germany"`
}

type personalSchema struct {
	Name               string          `json:"name,omitempty" schema:"description:full name;example:John Doe"`
	Gender             string          `json:"gender,omitempty" schema:"example:male"`
	Image              string          `json:"image,omitempty" schema:"description:url of the personal photo;example:example.com/Abcxyz"`
	Email              string          `json:"email,omitempty" schema:"description:email address;format:idn-email;example:lucas@example.com"`
	Phone              string          `json:"phone,omitempty" schema:"description:phone number;example:912-217-7923"`
	URL                string          `json:"url,omitempty" schema:"description:homepage url;format:uri;example:http://www.example.com/"`
	Summary            string          `json:"summary,omitempty" schema:"description:a short sentence about yourself;example:The man who sold the world!"`
	CurrentLocation    locationSchema  `json:"currentLocation,omitempty" schema:"description:living location"`
	PermanentLocation  locationSchema  `json:"permanentLocation,omitempty" schema:"description:permanently living location"`
	Birthday           string          `json:"birthday,omitempty" schema:"description:birthday date;format:date;example:1995-02-14"`
	BirthPlace         string          `json:"birthPlace,omitempty" schema:"description:place of birth;example:Munich"`
	Profiles           []profileSchema `json:"profiles,omitempty" schema:"description:list of social networks"`
	RelationshipStatus string          `json:"relationshipStatus,omitempty" schema:"description:civil status;example:married"`
	PostalAddress      string          `json:"postalAddress,omitempty" schema:"example:John Doe, Wittekindshof, Schulstrasse 4, 32547 Bad Oyenhausen, Germany"`
}

type workSchema struct {
	Name        string         `json:"name" schema:"description:name of company;example:XYZ Inc"`
	Description string         `json:"description,omitempty" schema:"description:companies primary activity;example:A social media company"`
	Position    string         `json:"position" schema:"description:position at the company;example:Software Engineer"`
	Location    locationSchema `json:"location,omitempty" schema:"description:location of the company"`
	URL         string         `json:"url,omitempty" schema:"description:url of the company website;format:uri;example:http://xyz.example.com"`
	StartDate   string         `json:"startDate" schema:"description:start date;format:date;example:2017-12-29"`
	EndDate     string         `json:"endDate,omitempty" schema:"description:end date;format:date;example:2018-12-29"`
	Summary     string         `json:"summary,omitempty" schema:"description:an overview of responsibilities;example:Developing and maintaining the company website using syna"`
	Highlights  []string       `json:"highlights,omitempty" schema:"description:some of accomplishments;items_example:Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"`
}

type educationSchema struct {
	Institution string         `json:"institution" schema:"name of institute;example:XYZ Institute of Technology"`
	Location    locationSchema `json:"location,omitempty" schema:"description:location of institution"`
	Area        string         `json:"area" schema:"description:area of study;example:Engineering"`
	StudyType   string         `json:"studyType" schema:"description:type of study;example:Bachelor"`
	StartDate   string         `json:"startDate" schema:"description:start date;format:date;example:2017-06-28"`
	EndDate     string         `json:"endDate,omitempty" schema:"description:end date;format:date;example:2013-06-28"`
	Score       scoreSchema    `json:"score,omitempty"`
	Courses     []string       `json:"courses,omitempty" schema:"description:notable courses/subjects;items_example:CS302 - Introduction to Algorithms"`
	Honors      []string       `json:"honors,omitempty" schema:"description:some education honours;items_example:Magna Cum Laude"`
	Highlights  []string       `json:"highlights,omitempty" schema:"description:some of accomplishments;items_example:Live abroad within a new culture"`
}

type volunteerSchema struct {
	Organization string         `json:"organization" schema:"description:name of the organization;example:Xyz"`
	Position     string         `json:"position" schema:"description:type of the contribution;example:Open Source Contributor"`
	Location     locationSchema `json:"location,omitempty" schema:"description:location of activity"`
	URL          string         `json:"url,omitempty" schema:"description:related link to support volunteer experience;format:uri;example:http://xyz.example.com"`
	StartDate    string         `json:"startDate" schema:"description:start date;format:date;example:2014-06-29"`
	EndDate      string         `json:"endDate,omitempty" schema:"description:end date;format:date;example:2017-06-29"`
	Summary      string         `json:"summary,omitempty" schema:"description:an overview of responsibilities;example:Frontend developer"`
	Highlights   []string       `json:"highlights,omitempty" schema:"description:some of accomplishments;items_example:Invited as a speaker in Xyzcon'17"`
}

type publicationSchema struct {
	Name            string           `json:"name" schema:"description:name of the publication;example:Deep learning and Artificial Intelligence"`
	Publisher       string           `json:"publisher" schema:"description:name of the publisher;example:XYZ, Computer Magazine"`
	PublicationDate string           `json:"publicationDate,omitempty" schema:"description:date of publication;format:date;example:2015-08-01"`
	Resources       []resourceSchema `json:"resources,omitempty" schema:"description:multiple resources with label"`
	URL             string           `json:"url,omitempty" schema:"description:url of the publication;format:uri;example:http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html"`
	Summary         string           `json:"summary,omitempty" schema:"description:short summary of the publication;example:Discussion of the advent of deep learning and artificial intelligence"`
}

type legalSchema struct {
	Name             string           `json:"name" schema:"description:name of the document;example:XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format"`
	LegalType        string           `json:"legalType,omitempty" schema:"description:type of the document;example:Patent, Trademark, Copyright"`
	Description      string           `json:"description,omitempty" schema:"description:a brief description about the document;example:Some legal document!"`
	ApplicationDate  string           `json:"applicationDate,omitempty" schema:"description:date of the application;format:date;example:2015-08-01"`
	GrantDate        string           `json:"grantDate,omitempty" schema:"description:date of the grant;format:date;example:2016-09-01"`
	EndDate          string           `json:"endDate,omitempty" schema:"description:end date;format:date;example:2020-09-03"`
	Resources        []resourceSchema `json:"resources,omitempty" schema:"description:multiple resources with label"`
	IDNumber         string           `json:"idNumber,omitempty" schema:"description:application number or Id Number;example:JP2004369746A"`
	CurrentAssignee  string           `json:"currentAssignee,omitempty" schema:"example:John Doe"`
	PreviousAssignee string           `json:"previousAssignee,omitempty" schema:"example:John Doe"`
	Author           string           `json:"author,omitempty" schema:"example:John Doe"`
	CoAuthors        string           `json:"coAuthors,omitempty" schema:"example:John Doe"`
}

type skillSchema struct {
	Name        string   `json:"name" schema:"description:name of the skill;example:Web Development"`
	Proficiency string   `json:"proficiency" schema:"description:proficiency level of the skill;enum:beginner,early,competent,advanced,expert;example:advanced"`
	Keywords    []string `json:"keywords,omitempty" schema:"description:some keywords pertaining to the skill;items_example:HTML"`
}

type awardSchema struct {
	Title   string `json:"title" schema:"description:title of the award;example:Awarded Software Process Achievement Award"`
	Date    string `json:"date" schema:"description:date of the award;format:date;example:2016-06-12"`
	Awarder string `json:"awarder" schema:"description:name of the awarder;example:IEEE"`
	Summary string `json:"summary,omitempty" schema:"description:reason of the award;example:Received for my work in Deep learning and AI"`
}

type projectSchema struct {
	Name        string           `json:"name" schema:"description:name of the project;example:File Transfer application"`
	Location    locationSchema   `json:"location,omitempty" schema:"description:location of the project"`
	Description string           `json:"description,omitempty" schema:"description:short summary of project;example:Developed a client and server based application"`
	Highlights  []string         `json:"highlights,omitempty" schema:"description:specify multiple features;items_example:used Java AWT and Swing for client side userinterface"`
	Keywords    []string         `json:"keywords,omitempty" schema:"description:specify special elements involved;items_example:Java"`
	StartDate   string           `json:"startDate" schema:"start date;format:date;example:2016-06-29"`
	EndDate     string           `json:"endDate,omitempty" schema:"description:end date;format:date;example:2017-03-02"`
	Resources   []resourceSchema `json:"resources,omitempty" schema:"description:specify multiple resources with label"`
	URL         string           `json:"url,omitempty" schema:"description:url of the project;format:uri;example:http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html"`
	Roles       []string         `json:"roles,omitempty" schema:"description:specify your role on this project;items_example:Team Lead"`
	Entity      string           `json:"entity,omitempty" schema:"description:relevant company/entity affiliations;example:greenpeace"`
	Type        string           `json:"type,omitempty" schema:"description:type of the project;example:volunteering"`
}

type certificateSchema struct {
	Code          string      `json:"code,omitempty" schema:"description:code of the certificate;example:1Z0-062"`
	Name          string      `json:"name" schema:"description:name of the certificate;example:XYZ Certified Application Specialist (MCAS)"`
	Website       string      `json:"website,omitempty" schema:"description:link to issuing authority's description of the certificate;format:uri;example:http://www.example.org"`
	Verification  string      `json:"verification,omitempty" schema:"description:external candidate verification URL;format:uri;example:http://www.example.org"`
	GrantDate     string      `json:"grantDate,omitempty" schema:"description:date of the grant;format:date;example:2017-06-29"`
	Score         scoreSchema `json:"score,omitempty" schema:"description:exam score"`
	EndDate       string      `json:"endDate,omitempty" schema:"description:end date of certificate;format:date;example:2017-06-29"`
	DoesNotExpire bool        `json:"doesNotExpire,omitempty" schema:"example:true"`
}

type referenceSchema struct {
	Name      string `json:"name,omitempty" schema:"description:name of the reference;example:Stephan Mark"`
	Company   string `json:"company,omitempty" schema:"description:company name;example:Xyz"`
	Position  string `json:"position,omitempty" schema:"description:position of reference;example:Senior Software Engineer"`
	Reference string `json:"reference" schema:"description:reference text;example:Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."`
}

type languageSchema struct {
	Language string `json:"language" schema:"description:name of language;example:English"`
	Level    string `json:"level" schema:"description:proficiency level for the language;enum:basic,conversational,fluent,native;example:fluent"`
}

type interestSchema struct {
	Name     string   `json:"name" schema:"description:name of the interest;example:Machine Learning"`
	Keywords []string `json:"keywords,omitempty" schema:"description:keywords of the interest;items_example:Neural Networks"`
}

type metaSchema struct {
	Canonical    string `json:"canonical,omitempty" schema:"description:URL (as per RFC 3986) to latest version of this document"`
	Version      string `json:"version,omitempty" schema:"description:version field which follows semver;example:v1.0.0"`
	LastModified string `json:"lastModified,omitempty" schema:"description:date-time of last modified;format:date-time;example:2017-06-29T15:53:01+01:00"`
	UUID         string `json:"uuid,omitempty" schema:"description:uuid v4 of the resume;example:078c39ce-23ee-4970-9637-c07379132dce"`
}

// Schema is the struct which jsonschema, uischema, and example are generated from.
// This package will use `schema` tags to generate this information. Also, it
// is important to note the `ui:order` in uischema is generated based on the
// properties order in the struct.
type Schema struct {
	Core         coreSchema          `json:"core"`
	Personal     personalSchema      `json:"personal" schema:"description:sensitive informations"`
	Work         []workSchema        `json:"work"`
	Education    []educationSchema   `json:"education"`
	Volunteer    []volunteerSchema   `json:"volunteer"`
	Publications []publicationSchema `json:"publications" schema:"description:list of publications"`
	Legal        []legalSchema       `json:"legal" schema:"description:list of legals"`
	Skills       []skillSchema       `json:"skills" schema:"description:list of professional skill-sets"`
	Awards       []awardSchema       `json:"awards" schema:"description:list of awards have been received"`
	Projects     []projectSchema     `json:"projects" schema:"description:list of career projects"`
	Certificates []certificateSchema `json:"certificate"`
	References   []referenceSchema   `json:"references" schema:"description:list of references"`
	Languages    []languageSchema    `json:"languages" schema:"description:list of languages"`
	Interests    []interestSchema    `json:"interests"`
	Meta         metaSchema          `json:"meta" schema:"description:the schema version and any other tooling configuration"`
}
