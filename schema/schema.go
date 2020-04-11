package schema

type locationSchema struct {
	Lat  float64 `json:"lat" schema:"description:latitude part of a coordinate;example:48.1548895"`
	Long float64 `json:"long" schema:"description:longitude part of a coordinate;example:11.4717966"`
}

type areaSchema struct {
	Location locationSchema `json:"location" schema:"description:specific location"`
	Radius   string         `json:"radius" schema:"description:radius denoting the area around a location;example:5km"`
}

type scoreSchema struct {
	Value string `json:"value" schema:"description:value of the score;example:3.4"`
	Best  string `json:"best" schema:"description:best possible score;example:4"`
	Worst string `json:"worst" schema:"description:worst possible score;example:0"`
	Label string `json:"label,omitempty" schema:"description:label/type of the score;example:GPA"`
}

type resourceSchema struct {
	URL   string `json:"url" schema:"description:url of the resource;format:uri;example:http://www.example.com/my-example-slides/"`
	Label string `json:"label" schema:"description:label of the resource;example:Slides"`
}

type profileSchema struct {
	Network  string `json:"network" schema:"description:name of the network;example:Github"`
	Username string `json:"username" schema:"description:username in the network;example:john_doe"`
	URL      string `json:"url,omitempty" schema:"description:url of the network profile;format:uri;example:https://github.com/john_doe"`
}

type peopleSchema struct {
	Name string `json:"name" schema:"description:full name;example:John Doe"`
	URL  string `json:"url,omitempty" schema:"description:url to link to mentioned person;format:uri;example:https://linkedin.com/john_doe"`
}

type idSchema struct {
	ID    string `json:"id" schema:"description:ID/code of the identifier;example:10.1000/182"`
	Label string `json:"label" schema:"description:label of the identifier;example:DOI"`
}

type coreSchema struct {
	Title       string            `json:"title,omitempty" schema:"description:job title;example:Software Engineer"`
	LivingArea  areaSchema        `json:"livingArea,omitempty" schema:"description:living area which could be city or country or even continent"`
	Disposition dispositionSchema `json:"disposition,omitempty" schema:"description: disposition or desire for work engagements"`
	Summary     string            `json:"summary,omitempty" schema:"description:write a short biography about you and your achievements optionally in relation to your current application;example:Passionate individual looking for my next xyz challenge"`
}

type dispositionSchema struct {
	TravelTime     string           `json:"travelTime,omitempty" schema:"description:percentage of time willing to travel;example:25"`
	Authorizations []string         `json:"authorization,omitempty" schema:"description:authorizations such as work and travel visa and residencies;items_example:US Permanent Resident Card"`
	CommitmentType string           `json:"commitmentType,omitempty" schema:"description:types of work commitements desired such as fulltime, part-time, contract, permanent etc.;items_example:contract"`
	Remote         bool             `json:"remote,omitempty" schema:"description:desire to work remotely;example:true"`
	WorkArea       areaSchema       `json:"workArea,omitempty" schema:"description:work area which could be city or country or even remote"`
	Relocation     relocationSchema `json:"relocation,omitempty" schema:"description:relocation preferences such as willingness and location"`
}

type relocationSchema struct {
	Willingness bool   `json:"willingness,omitempty" schema:"description:willingness to relocate;example:true"`
	Area        string `json:"area,omitempty" schema:"description:work area which could be city or country or language area;example:English Speaking Countries"`
}

type entitySchema struct {
	Name        string `json:"name" schema:"name of the institute, company or such;example:Entity of Things"`
	Description string `json:"description,omitempty" schema:"description:entity's primary activity;example:A good things entity"`
	URL         string `json:"url,omitempty" schema:"description:url of the entity website;format:uri;example:http://xyz.example.com"`
	Image       string `json:"image,omitempty" schema:"description:url or src of an image, photo or logo;example:https://resumic.org/assets/entity.svg"`
}

type personalSchema struct {
	Name               string          `json:"name,omitempty" schema:"description:full name;example:John Doe"`
	LegalGender        string          `json:"legalGender,omitempty" schema:"description:legal gender;example:male"`
	IdentityGender     string          `json:"identityGender,omitempty" schema:"description:gender you are optionally identifying with;example:male"`
	Pronouns           string          `json:"pronouns,omitempty" schema:"description:pronouns you feel comfortable with being used;example:they/their/them"`
	Email              string          `json:"email,omitempty" schema:"description:email address;format:idn-email;example:lucas@example.com"`
	Phone              string          `json:"phone,omitempty" schema:"description:phone number;example:912-217-7923"`
	CurrentLocation    locationSchema  `json:"currentLocation,omitempty" schema:"description:current living location"`
	PermanentLocation  locationSchema  `json:"permanentLocation,omitempty" schema:"description:permanently living location"`
	Birthday           string          `json:"birthday,omitempty" schema:"description:date of birthday;format:date;example:1995-02-14"`
	BirthPlace         string          `json:"birthPlace,omitempty" schema:"description:place of birth;example:Munich"`
	Citizenship        string          `json:"citizenship,omitempty" schema:"description:country of citizenship;example:Germany"`
	Profiles           []profileSchema `json:"profiles,omitempty" schema:"description:list of social/other networks"`
	RelationshipStatus string          `json:"relationshipStatus,omitempty" schema:"description:relevant relationship or civil status;example:married"`
	PostalAddress      string          `json:"postalAddress,omitempty" schema:"description:full address useable to send postal mail;example:John Doe, Wittekindshof, Schulstrasse 4, 32547 Bad Oyenhausen, Germany"`
	URL                string          `json:"url,omitempty" schema:"description:homepage url;format:uri;example:http://www.example.com/"`
	Image              string          `json:"image,omitempty" schema:"description:url or src of an image, photo or logo;example:https://resumic.org/assets/profile.svg"`
}

type workSchema struct {
	Work    []workItemSchema `json:"work,omitempty" schema:"description:list of work experiences"`
	Summary string           `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full work section;example:My general work experience was interrupted by a new educational experience, which now enables me to do even more."`
}

type workItemSchema struct {
	Roles      []string       `json:"roles" schema:"description:concurrent roles/positions at the company;items_example:Software Engineer"`
	Entity     entitySchema   `json:"entity" schema:"description:entity details"`
	StartDate  string         `json:"startDate" schema:"description:start date;format:date;example:2017-12-29"`
	EndDate    string         `json:"endDate,omitempty" schema:"description:end date;format:date;example:2018-12-29"`
	Location   locationSchema `json:"location,omitempty" schema:"description:location of the activity"`
	Summary    string         `json:"summary,omitempty" schema:"description:continous text summarizing the work, time and experiences;example:The work with company XYZ was a great experience taking my knowledge accumulation to new heights. The remote team setup made me aware of my love for remote working."`
	Highlights []string       `json:"highlights,omitempty" schema:"description:important (bullet point) achievements;items_example:Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers"`
}

type educationSchema struct {
	Education []educationItemSchema `json:"education,omitempty" schema:"description:list of education experiences"`
	Summary   string                `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full education section;example:My diverse education gives me an additional context in my area of work."`
}

type educationItemSchema struct {
	Degree     string         `json:"degree,omitempty" schema:"description:final or persuing degree;example:Bachelor of Science"`
	Area       string         `json:"area,omitempty" schema:"description:area of study;example:Engineering"`
	Entity     entitySchema   `json:"entity" schema:"description:entity details"`
	StudyStage string         `json:"studyStage" schema:"description:stage of study;enum:primary,secondary,higher;example:higher"`
	StartDate  string         `json:"startDate" schema:"description:start date;format:date;example:2017-06-28"`
	EndDate    string         `json:"endDate,omitempty" schema:"description:end date;format:date;example:2013-06-28"`
	Location   locationSchema `json:"location,omitempty" schema:"description:location of the activity"`
	Score      scoreSchema    `json:"score,omitempty"`
	Honors     []string       `json:"honors,omitempty" schema:"description:specific honors;items_example:Magna Cum Laude"`
	Courses    []string       `json:"courses,omitempty" schema:"description:notable courses/subjects;items_example:CS302 - Introduction to Algorithms"`
	Summary    string         `json:"summary,omitempty" schema:"description:continous text summarizing the work, time and experiences;example:My studies with institute XYZ was a great experience taking my knowledge accumulation to new heights. The campus live not only enabled meeting new friends, but to build out my network."`
	Highlights []string       `json:"highlights,omitempty" schema:"description:important (bullet point) achievements;items_example:Live abroad within a new culture"`
}

type volunteerSchema struct {
	Volunteer []volunteerItemSchema `json:"volunteer,omitempty" schema:"description:list of volunteering experiences"`
	Summary   string                `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full volunteer section;example:Creating positive change is not only important in my professional life, but also in my private life and I do my best to help and volunteer."`
}

type volunteerItemSchema struct {
	Roles      []string         `json:"roles" schema:"description:roles/positions or type of the contribution;items_example:Open Source Contributor"`
	Entity     entitySchema     `json:"entity" schema:"description:entity details"`
	StartDate  string           `json:"startDate" schema:"description:start date;format:date;example:2014-06-29"`
	EndDate    string           `json:"endDate,omitempty" schema:"description:end date;format:date;example:2017-06-29"`
	Location   locationSchema   `json:"location,omitempty" schema:"description:location of the activity"`
	Resources  []resourceSchema `json:"resources,omitempty" schema:"description:additional linkable resources identified with a label"`
	Summary    string           `json:"summary,omitempty" schema:"description:continous text summarizing the experiences;example:Doing accounting work to keep things afloat."`
	Highlights []string         `json:"highlights,omitempty" schema:"description:important (bullet point) achievements;items_example:Set up huge donor run"`
}

type publicationSchema struct {
	Publications []publicationItemSchema `json:"publications,omitempty" schema:"description:list of publications"`
	Summary      string                  `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full publication section;example:Thanks to my co-Authors and the willing reviewers my publications far surpass my imaginations."`
}

type publicationItemSchema struct {
	Name            string           `json:"name" schema:"description:name of the publication;example:Deep learning and Artificial Intelligence"`
	Description     string           `json:"description,omitempty" schema:"description:short description about the publication;example:Discussion of the advent of deep learning and artificial intelligence"`
	Entity          entitySchema     `json:"entity" schema:"description:entity details"`
	PublicationDate string           `json:"publicationDate,omitempty" schema:"description:date of publication;format:date;example:2015-08-01"`
	Roles           []string         `json:"roles,omitempty" schema:"description:specific roles within or around this publication - author might be assumed as default;items_example:Reviewer"`
	CoAuthors       []peopleSchema   `json:"coAuthors,omitempty" schema:"description:co-authors for this publication"`
	Identifier      []idSchema       `json:"identifier,omitempty" schema:"description:labeled identifiers"`
	URL             string           `json:"url,omitempty" schema:"description:url of the publication;format:uri;example:http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html"`
	Image           string           `json:"image,omitempty" schema:"description:url or src of an image, photo or logo;example:https://resumic.org/assets/publication.svg"`
	Resources       []resourceSchema `json:"resources,omitempty" schema:"description:additional linkable resources identified with a label"`
	Summary         string           `json:"summary,omitempty" schema:"description:continous text summarizing the experiences;example:Writing this publication opened up a huge wave of new research areas. The work together with <person> was the highlight of my research in the field of AI."`
	Highlights      []string         `json:"highlights,omitempty" schema:"description:important (bullet point) achievements;items_example:Changed how we see the world"`
}

type legalSchema struct {
	Legal   []legalItemSchema `json:"legal,omitempty" schema:"description:list of legal"`
	Summary string            `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full legal section;example:As a hobby inventor I was able to amass a wide variety of copyright, patent and other legal objects."`
}

type legalItemSchema struct {
	Name             string           `json:"name" schema:"description:name of the document, patent, copyright assignment or the likes;example:XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format"`
	Description      string           `json:"description,omitempty" schema:"description:short description about the document;example:Some legal document!"`
	Entity           entitySchema     `json:"entity" schema:"description:entity details"`
	LegalType        string           `json:"legalType,omitempty" schema:"description:type of the document;example:Patent, Trademark, Copyright"`
	ApplicationDate  string           `json:"applicationDate,omitempty" schema:"description:date of the application;format:date;example:2015-08-01"`
	GrantDate        string           `json:"grantDate,omitempty" schema:"description:date of the grant;format:date;example:2016-09-01"`
	EndDate          string           `json:"endDate,omitempty" schema:"description:end date;format:date;example:2020-09-03"`
	Resources        []resourceSchema `json:"resources,omitempty" schema:"description:multiple resources with label"`
	Identifier       []idSchema       `json:"identifier,omitempty" schema:"description:labeled identifiers"`
	CurrentAssignee  string           `json:"currentAssignee,omitempty" schema:"description:current entity or individual assigned to the legal item;example:John Doe"`
	PreviousAssignee string           `json:"previousAssignee,omitempty" schema:"description:previous entity or individual assigned to the legal item;example:John Doe"`
	Roles            []string         `json:"roles,omitempty" schema:"description:specific roles within or around this legal item - author might be assumed as default;items_example:Reviewer"`
	CoAuthors        []peopleSchema   `json:"coAuthors,omitempty" schema:"description:co-authors for this legal item"`
	Summary          string           `json:"summary,omitempty" schema:"description:summary of achievements;example:Build out onboarding process as core maintainer"`
}

type skillSchema struct {
	Skills  []skillItemSchema `json:"skills,omitempty" schema:"description:list of skills"`
	Summary string            `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full skill section;example:My passion for figuring out things brings me to discover and learn new skills. I only listed only the most relevant skills. Please reach out for any further questions."`
}

type skillItemSchema struct {
	Name        string   `json:"name" schema:"description:name of the skill;example:Web Development"`
	Proficiency string   `json:"proficiency" schema:"description:proficiency level of the skill using numerical declaration 1=beginner, 2=early, 3=competent, 4=advanced, 5=expert;enum:1,2,3,4,5;example:4"`
	Keywords    []string `json:"keywords,omitempty" schema:"description:keywords giving context or details about the skill;items_example:HTML"`
}

type awardSchema struct {
	Awards  []awardItemSchema `json:"awards,omitempty" schema:"description:list of awards"`
	Summary string            `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full award section;example:Most awards are of honory mention, but I personally want to point out Award XYZ as it is one of my most precious ones."`
}

type awardItemSchema struct {
	Name        string       `json:"name" schema:"description:name of the award;example:Nobel Peace Prize"`
	Description string       `json:"description,omitempty" schema:"description:short description about the award;example:slightly known award for supposedly doing the good fight"`
	Date        string       `json:"date" schema:"description:date of the award;format:date;example:2016-06-12"`
	Entity      entitySchema `json:"entity" schema:"description:entity details;"`
	Summary     string       `json:"summary,omitempty" schema:"description:continous text summarizing the experiences;example:Due to my long time work on preventing harm, I was lucky to be awarded the Nobel Peace Prize."`
}

type projectSchema struct {
	Projects []projectItemSchema `json:"projects,omitempty" schema:"description:list of projects"`
	Summary  string              `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full project section;example:Most of my projects are work related and can't be shown extensively. I'm happy to talk about them in detail on a call."`
}

type projectItemSchema struct {
	Name        string           `json:"name" schema:"description:name of the project, contribution, talk or the like;example: Resumic"`
	Description string           `json:"description,omitempty" schema:"description:short description of the project;example:Open CV standard"`
	StartDate   string           `json:"startDate" schema:"description:start date;format:date;example:2016-06-29"`
	EndDate     string           `json:"endDate,omitempty" schema:"description:end date;format:date;example:2017-03-02"`
	Roles       []string         `json:"roles,omitempty" schema:"description:specific roles within or around this project;items_example:Team Lead"`
	Entity      entitySchema     `json:"entity" schema:"description:entity details"`
	Type        string           `json:"type,omitempty" schema:"description:type of the project;example:open source|volunteering|hobby"`
	Location    locationSchema   `json:"location,omitempty" schema:"description:location of the project"`
	URL         string           `json:"url,omitempty" schema:"description:url of the project;format:uri;example:http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html"`
	Image       string           `json:"image,omitempty" schema:"description:url or src of an image, photo or logo;example:https://resumic.org/assets/project.svg"`
	Resources   []resourceSchema `json:"resources,omitempty" schema:"description:additional linkable resources identified with a label"`
	Summary     string           `json:"summary,omitempty" schema:"description:continous text summarizing the experiences;example:Starting a new CV standard to create a new CV was definitely not the most efficient use of time."`
	Highlights  []string         `json:"highlights,omitempty" schema:"description:important (bullet point) achievements;items_example:Used by a broad community"`
	Keywords    []string         `json:"keywords,omitempty" schema:"description:keywords giving context or details about the project;items_example:JSON"`
}

type certificateSchema struct {
	Certificates []certificateItemSchema `json:"certificates,omitempty" schema:"description:list of certificates"`
	Summary      string                  `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full certificate section;example:Personally prefer the technology over specific certificates, but happy to take additional certification exams to validate my experiences."`
}

type certificateItemSchema struct {
	Name            string       `json:"name" schema:"description:name of the certificate;example:XYZ Certified Application Specialist (MCAS)"`
	Description     string       `json:"description,omitempty" schema:"description:short description of the certificate;example:software application specialization certificate"`
	Entity          entitySchema `json:"entity" schema:"description:entity details"`
	StartDate       string       `json:"startDate,omitempty" schema:"description:start date;format:date;example:2017-06-29"`
	EndDate         string       `json:"endDate,omitempty" schema:"description:end date;format:date;example:2017-06-29"`
	Code            string       `json:"code,omitempty" schema:"description:code of the certificate;example:1Z0-062"`
	Score           scoreSchema  `json:"score,omitempty" schema:"description:exam score"`
	VerificationURL string       `json:"verification,omitempty" schema:"description:public certificate verification URL;format:uri;example:https://www.coursera.org/account/accomplishments/verify/XYZ"`
	Perpetual       string       `json:"perpetual,omitempty" schema:"description:validity status being perpetual or not;example:true"`
	Image           string       `json:"image,omitempty" schema:"description:url or src of a certification badge or another related image;example:https://resumic.org/assets/certificate.svg"`
}

type referenceSchema struct {
	References []referenceItemSchema `json:"references,omitempty" schema:"description:list of references"`
	Summary    string                `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full reference section;example:My references are great due to specific items. For contact details about my references please reach out directly"`
}

type referenceItemSchema struct {
	Name      string       `json:"name,omitempty" schema:"description:name of the reference;example:Stephan Mark"`
	Reference string       `json:"reference" schema:"description:reference text;example: Person was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."`
	Entity    entitySchema `json:"entity" schema:"description:entity details"`
	Role      string       `json:"role,omitempty" schema:"description:role/position of the reference at affiliated entity;example:Senior Software Engineer"`
}

type languageSchema struct {
	Languages []languageItemSchema `json:"languages,omitempty" schema:"description:list of languages"`
	Summary   string               `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full language section;example:Due to an extended time in a country my fluency is great."`
}

type languageItemSchema struct {
	Language string `json:"language" schema:"description:name of language;example:English"`
	Level    string `json:"level" schema:"description:proficiency level for the language using numerical declaration 1=basic, 2=conversational, 3=fluent, 4=native;enum:1,2,3,4;example:1"`
}

type interestSchema struct {
	Interests []interestItemSchema `json:"interests,omitempty" schema:"description:list of interests"`
	Summary   string               `json:"summary,omitempty" schema:"description:continous text summarizing or adding context to the full interest section;example:My interest give me the necessary downtime to balance work and family obligations."`
}

type interestItemSchema struct {
	Name     string   `json:"name" schema:"description:name of the interest;example:Machine Learning"`
	Keywords []string `json:"keywords,omitempty" schema:"description:keywords giving context or details about the project;items_example:Neural Networks"`
}

type metaSchema struct {
	Canonical     string           `json:"canonical,omitempty" schema:"description:URL (as per RFC 3986) to the latest version of this document;format:uri;example:http://example.org/my-resume.json"`
	Version       string           `json:"version,omitempty" schema:"description:version field which follows semver;example:v1.0.0"`
	LastModified  string           `json:"lastModified,omitempty" schema:"description:date-time of last modified;format:date-time;example:2017-06-29T15:53:01+01:00"`
	UUID          string           `json:"uuid,omitempty" schema:"description:uuid v4 of the resume;example:078c39ce-23ee-4970-9637-c07379132dce"`
	Tooling       struct{}         `json:"tooling,omitempty" schema:"description:empty section for tools to write to in the form of meta.tooling.<tool>.<tool-data-section>"`
	Schema        schemaToolSchema `json:"schema,omitempty" schema:"description:empty section for the schema itself to hold tooling relevant data"`
	SchemaVersion string           `json:"schemaVersion,omitempty" schema:"description:version field for used schema which follows semver;example:v1.0.0"`
}

type schemaToolSchema struct {
	Filters      struct{} `json:"filters,omitempty" schema:"description:empty section the schema itself to add filter data"`
	Localization struct{} `json:"localization,omitempty" schema:"description:empty section to add lolcalization data"`
}

// Schema is the struct which schema.json, ui-schema.json, and schema-example.json are generated from.
// This package will use `schema` tags to generate this information.
// It is important to note the `ui:order` in uischema is generated based on the
// properties order in the struct.
type Schema struct {
	Schema       string            `json:"$schema" schema:"description:url of JSON Schema used for validation;format:uri;example:https://schema.resumic.org/latest.json"`
	Core         coreSchema        `json:"core" schema:"description:core non identifying facts"`
	Personal     personalSchema    `json:"personal" schema:"description:sensitive and potentially identifying information - should not be requested or used"`
	Work         workSchema        `json:"work" schema:"description:language section with summary and list of work experiences centered on the entity - bigger projects could be showcased using the projects section"`
	Education    educationSchema   `json:"education" schema:"description:language section with summary and list of the education"`
	Volunteer    volunteerSchema   `json:"volunteer" schema:"description:language section with summary and list of volunteer experiences centered on the volunteer organisation"`
	Publications publicationSchema `json:"publications" schema:"description:language section with summary and list of written, spoken or conference publications"`
	Legal        legalSchema       `json:"legal" schema:"description:language section with summary and list of legal items"`
	Skills       skillSchema       `json:"skills" schema:"description:language section with summary and list of professional skill-sets"`
	Awards       awardSchema       `json:"awards" schema:"description:language section with summary and list of received awards"`
	Projects     projectSchema     `json:"projects" schema:"description:language section with summary and list of (career-) relevant projects"`
	Certificates certificateSchema `json:"certificates" schema:"description:language section with summary and list of certificates and permits"`
	References   referenceSchema   `json:"references" schema:"description:language section with summary and list of references"`
	Languages    languageSchema    `json:"languages" schema:"description:language section with summary and list of languages"`
	Interests    interestSchema    `json:"interests" schema:"description:language section with summary and list of interests"`
	Meta         metaSchema        `json:"meta" schema:"description: meta data such as resume version as well as tool provided and generated data/configuration"`
}
