package schema

type locationSchema struct {
	Lat  float64 `json:"lat" schema:"example:35.712758"`
	Long float64 `json:"long" schema:"example:51.392114"`
}

type scoreSchema struct {
	Type  string `json:"type" schema:"description:eg. GPA or PERCENTAGE - [Add score type];example:GPA"`
	Value string `json:"value" schema:"description:eg. 3.4/4.0 - [Add obtained score/total score];example:3.4"`
}

type resourceSchema struct {
	URL   string `json:"url" schema:"description:e.g. http://www.example.com/my-example-slides/;format:uri;example:http://www.example.com/my-example-slides/"`
	Label string `json:"label" schema:"description:e.g Slides;example:Slides"`
}

type coreSchema struct {
	Name              string         `json:"name" schema:"description:e.g. John Doe;example:John Doe"`
	Title             string         `json:"title" schema:"description:e.g. Software Engineer;example:Software Engineer"`
	Image             string         `json:"image" schema:"description:e.g. example.com/Abcxyz - [URL (as per RFC 3986) to a image in JPEG or PNG format];example:example.com/Abcxyz"`
	Email             string         `json:"email" schema:"description:e.g. lucas@example.com;format:email;example:lucas@example.com"`
	Phone             string         `json:"phone" schema:"description:e.g. 912-217-7923 - [Phone number];example:912-217-7923"`
	URL               string         `json:"url" schema:"description:e.g. http://www.example.com/my-slides/;format:uri;example:http://www.example.com/"`
	Summary           string         `json:"summary" schema:"description:Write a short 2-3 sentence biography about yourself;example:The man who sold the world!"`
	CurrentLocation   locationSchema `json:"currentLocation" schema:"description:Select the location where you currently live."`
	PermanentLocation locationSchema `json:"permanentLocation" schema:"description:Select the location where you permanently live."`
}

type workSchema struct {
	Name        string         `json:"name" schema:"description:e.g. XYZ Inc. - [Company name];example:XYZ Inc"`
	Description string         `json:"description" schema:"description:e.g. A social media company - [Description of the companies primary activity];example:A social media company"`
	Position    string         `json:"position" schema:"description:e.g. Software Engineer - [Position at the company];example:Software Engineer"`
	Location    locationSchema `json:"location" schema:"description:e.g. Germany - [Location for this activity]"`
	URL         string         `json:"url" schema:"description:e.g. http://xyz.example.com - [Related link to the company website];format:uri;example:http://xyz.example.com"`
	StartDate   string         `json:"startDate" schema:"description:e.g. 2017-06-28 - [resume.json uses the ISO 8601 date standard];format:date;example:2017-12-29"`
	EndDate     string         `json:"endDate" schema:"description:e.g. 2018-12-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2018-12-29"`
	Summary     string         `json:"summary" schema:"description:Give an overview of your responsibilities at the company;example:Developing and maintaining the company website using syna"`
	Highlights  []string       `json:"highlights" schema:"description:Specify multiple accomplishments;items_description:e.g. Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers;items_example:Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"`
}

type educationSchema struct {
	Institution string         `json:"institution" schema:"description:e.g. XYZ Institute of Technology - [Add institute name];example:XYZ Institute of Technology"`
	Location    locationSchema `json:"location" schema:"description:e.g. Germany - [Location for this institution]"`
	Area        string         `json:"area" schema:"description:e.g. Engineering;example:Engineering"`
	StudyType   string         `json:"studyType" schema:"description:e.g. Bachelor;example:Bachelor"`
	StartDate   string         `json:"startDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2017-06-28"`
	EndDate     string         `json:"endDate" schema:"description:e.g. 2013-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2013-06-28"`
	Score       scoreSchema    `json:"score"`
	Courses     []string       `json:"courses" schema:"description:List notable courses/subjects;items_description:e.g. CS302 - Introduction to Algorithms - [Add course name];items_example:CS302 - Introduction to Algorithms"`
	Honors      []string       `json:"honors" schema:"description:List education honours;items_description:e.g. Magna Cum Laude;items_example:Magna Cum Laude"`
}

type volunteerSchema struct {
	Organization string         `json:"organization" schema:"description:e.g. Xyz ;example:Xyz"`
	Position     string         `json:"position" schema:"description:e.g. Open Source Contributor - [Contribution type];example:Open Source Contributor"`
	Location     locationSchema `json:"location" schema:"description:e.g. Germany - [Location for this activity]"`
	URL          string         `json:"url" schema:"description:e.g. http://xyz.example.com - [Related link to support volunteer experience];format:uri;example:http://xyz.example.com"`
	StartDate    string         `json:"startDate" schema:"description:e.g. 2014-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2014-06-29"`
	EndDate      string         `json:"endDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ;format:date;example:2017-06-29"`
	Summary      string         `json:"summary" schema:"description:Give an overview of your responsibilities at the company;example:Frontend developer"`
	Highlights   []string       `json:"highlights" schema:"description:Specify accomplishments and achievements;items_description:e.g Invited as a speaker in Xyzcon'17;items_example:Invited as a speaker in Xyzcon'17"`
}

type publicationSchema struct {
	Name        string           `json:"name" schema:"description:e.g. Deep learning and Artificial Intelligence;example:Deep learning and Artificial Intelligence"`
	Publisher   string           `json:"publisher" schema:"description:e.g. XYZ, Computer Magazine;example:XYZ, Computer Magazine"`
	ReleaseDate string           `json:"releaseDate" schema:"description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard];example:2015-08-01"`
	Resources   []resourceSchema `json:"resources" schema:"description:Specify multiple resources with label"`
	URL         string           `json:"url" schema:"description:e.g. http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html;format:uri;example:http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html"`
	Summary     string           `json:"summary" schema:"description:e.g. Discussion of the advent of deep learning and artificial intelligence - [Short summary of publication];example:Discussion of the advent of deep learning and artificial intelligence"`
}

type legalSchema struct {
	Name            string           `json:"name" schema:"description:e.g. XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format - [Add document name];example:XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format"`
	LegalType       string           `json:"legalType" schema:"description:e.g. Patent, Trademark, Copyright - [Give the type of this document];example:Patent, Trademark, Copyright"`
	Description     string           `json:"description" schema:"description:Give a brief description about this document;example:Some legal document!"`
	ApplicationDate string           `json:"applicationDate" schema:"description:e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard];format:date;example:2015-08-01"`
	GrantDate       string           `json:"grantDate" schema:"description:e.g. 2016-09-01 - [resume.json uses the ISO 8601 date standard];format:date;example:2016-09-01"`
	EndDate         string           `json:"endDate" schema:"description:e.g. 2020-09-03 - [resume.json uses the ISO 8601 date standard];format:date;example:2020-09-03"`
	Resources       []resourceSchema `json:"resources" schema:"description:Specify multiple resources with label"`
	IDNumber        string           `json:"idNumber" schema:"description:e.g. JP2004369746A - [Add the application number or Id Number]  ;example:JP2004369746A"`
}

type skillSchema struct {
	Name    string      `json:"name" schema:"description:e.g. Web Development;example:Web Development"`
	Score   scoreSchema `json:"score" schema:"description:Score for the skill name"`
	Keyword []string    `json:"keyword" schema:"description:List some keywords pertaining to the skill;items_example:HTML"`
}

type awardSchema struct {
	Title   string `json:"title" schema:"description:e.g. Awarded Software Process Achievement Award ;example:Awarded Software Process Achievement Award"`
	Date    string `json:"date" schema:"description:e.g. 2016-06-12 - [resume.json uses the ISO 8601 date standard];format:date;example:2016-06-12"`
	Awarder string `json:"awarder" schema:"description:e.g. IEEE;example:IEEE"`
	Summary string `json:"summary" schema:"description:e.g. Received for my work in Deep learning and AI;example:Received for my work in Deep learning and AI"`
}

type projectSchema struct {
	Name        string           `json:"name" schema:"description:e.g. File Transfer application - [Name of the project];example:File Transfer application"`
	Location    locationSchema   `json:"location" schema:"description:e.g France - [Location for this activity]"`
	Description string           `json:"description" schema:"description:e.g. Developed a client and server based application - [Short summary of project];example:Developed a client and server based application"`
	Highlights  []string         `json:"highlights" schema:"description:Specify multiple features;items_description:e.g. used Java AWT and Swing for client side userinterface;items_example:used Java AWT and Swing for client side userinterface"`
	Keywords    []string         `json:"keywords" schema:"description:Specify special elements involved;items_description:e.g. Java;items_example:Java"`
	StartDate   string           `json:"startDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2016-06-29"`
	EndDate     string           `json:"endDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ;format:date;example:2017-03-02"`
	Resources   []resourceSchema `json:"resources" schema:"description:Specify multiple resources with label"`
	URL         string           `json:"url" schema:"description:e.g. http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html;format:uri;example:http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html"`
	Roles       []string         `json:"roles" schema:"description:Specify your role on this project or in company;items_description:e.g. Team Lead, Speaker, Writer;items_example:Team Lead"`
	Entity      string           `json:"entity" schema:"description:e.g. 'greenpeace', 'corporationXYZ' - [Relevant company/entity affiliations];example:greenpeace"`
	Type        string           `json:"type" schema:"description:e.g. 'volunteering', 'presentation', 'talk', 'application', 'conference';example:volunteering"`
}

type certificateSchema struct {
	Code          string      `json:"code" schema:"description:e.g. 1Z0-062;example:1Z0-062"`
	Name          string      `json:"name" schema:"description:e.g. XYZ Certified Application Specialist (MCAS) - [Add the certificate name];example:XYZ Certified Application Specialist (MCAS)"`
	Website       string      `json:"website" schema:"description:Link to issuing authority's description of the certificate;format:uri;example:http://www.example.org"`
	Verification  string      `json:"verification" schema:"description:External candidate verification URL;format:uri;example:http://www.example.org"`
	GrantDate     string      `json:"grantDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2017-06-29"`
	Score         scoreSchema `json:"score" schema:"description:Exam result (PASS/FAIL, 100%, 100)"`
	EndDate       string      `json:"endDate" schema:"description:e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard];format:date;example:2017-06-29"`
	DoesNotExpire bool        `json:"doesNotExpire" schema:"example:true"`
}

type referenceSchema struct {
	Name      string `json:"name" schema:"description:e.g. Stephan Mark;example:Stephan Mark"`
	Company   string `json:"company" schema:"description:e.g. Xyz;example:Xyz"`
	Position  string `json:"position" schema:"description:e.g. Senior Software Engineer;example:Senior Software Engineer"`
	Reference string `json:"reference" schema:"description:e.g. Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing.;example:Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."`
}

type languageSchema struct {
	Language string      `json:"language" schema:"description:e.g. English - [Name of language];example:English"`
	Score    scoreSchema `json:"score" schema:"description:Score for the language"`
}

type interestSchema struct {
	Name     string   `json:"name" schema:"description:e.g. Machine Learning;example:Machine Learning"`
	Keywords []string `json:"keywords" schema:"description:Specify keywords associated with the particular interest;items_description:e.g. Neural Networks, Convoluted Neural Networks;items_example:Neural Networks"`
}

type metaSchema struct {
	Canonical    string `json:"canonical" schema:"description:URL (as per RFC 3986) to latest version of this document"`
	Version      string `json:"version" schema:"description:e.g. v1.0.0 - [A version field which follows semver];example:v1.0.0"`
	LastModified string `json:"lastModified" schema:"description:e.g. 2017-06-29T15:53:00 - [resume.json uses the ISO 8601 date standard];format:date-time;example:2017-06-29T15:53:01+01:00"`
}

// Schema is the struct which jsonschema, uischema, and example are generated from.
// This package will use `schema` tags to generate this information. Also, it
// is important to note the `ui:order` in uischema is generated based on the
// properties order in the struct.
type Schema struct {
	Core         coreSchema          `json:"core"`
	Work         []workSchema        `json:"work"`
	Education    []educationSchema   `json:"education"`
	Volunteer    []volunteerSchema   `json:"volunteer"`
	Publications []publicationSchema `json:"publications" schema:"description:Specify your publications"`
	Legal        []legalSchema       `json:"legal" schema:"description:Specify your labels"`
	Skills       []skillSchema       `json:"skills" schema:"description:List out your professional skill-set"`
	Awards       []awardSchema       `json:"awards" schema:"description:Specify any awards you have received throughout your professional career"`
	Projects     []projectSchema     `json:"projects" schema:"description:Specify career projects"`
	Certificates []certificateSchema `json:"certificate"`
	References   []referenceSchema   `json:"references" schema:"description:List references you have received"`
	Languages    []languageSchema    `json:"languages" schema:"description:List any other languages you speak"`
	Interests    []interestSchema    `json:"interests"`
	Meta         metaSchema          `json:"meta" schema:"description:The schema version and any other tooling configuration lives here"`
}
