
# Core
This section is of object type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalProperties
##   properties
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
          "type": "string"
        }

###        title
Sub-section of type string, used to specify the title of the person
The schema snippet can be shown below:

       title:
        {
          "type": "string",
          "description": "e.g. Software Engineer"
        }

###        image
Sub-section of type string, used to specify the image of the person
The schema snippet can be shown below:

       image:
        {
          "type": "string",
          "description": "URL (as per RFC 3986) to a image in JPEG or PNG format"
        }

###        email
Sub-section of type string, used to specify the email of the person
The schema snippet can be shown below:

       email:
        {
          "type": "string",
          "description": "e.g. lucas@xyz.com",
          "format": "email"
        }

###        phone
Sub-section of type string, used to specify the phone of the person
The schema snippet can be shown below:

       phone:
        {
          "type": "string",
          "description": "e.g. 912-217-7923 [Phone numbers are stored as strings so use any format you like]"
        }

###        url
Sub-section of type string, used to specify the url of the person
The schema snippet can be shown below:

       url:
        {
          "type": "string",
          "format": "uri",
          "description": "e.g. http://www.exampleslide.com/my-slides/"
        }

###        summary
Sub-section of type string, used to specify the summary of the person
The schema snippet can be shown below:

       summary:
        {
          "type": "string",
          "description": "Write a short 2-3 sentence biography about yourself"
        }

###        currentLocation
Sub-section of type object, used to specify the currentLocation of the person
The schema snippet can be shown below:

       currentLocation:
        {
          "type": "object",
          "format": "location",
          "description": "Select the location where you currently live.",
          "properties": {
            "lat": {
              "type": "number"
            },
            "long": {
              "type": "number"
            }
          }
        }

###        permanentLocation
Sub-section of type object, used to specify the permanentLocation of the person
The schema snippet can be shown below:

       permanentLocation:
        {
          "type": "object",
          "format": "location",
          "description": "Select the location where you permanently live.",
          "properties": {
            "lat": {
              "type": "number"
            },
            "long": {
              "type": "number"
            }
          }
        }


# Work
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. XYZ Inc. - [Company name]"
          }

###        description
Sub-section of type string, used to specify the description of the person
The schema snippet can be shown below:

       description:
        {
            "type": "string",
            "description": "e.g. A social media company - [Description of the companies primary activity]"
          }

###        position
Sub-section of type string, used to specify the position of the person
The schema snippet can be shown below:

       position:
        {
            "type": "string",
            "description": "e.g. Software Engineer - [Position at the company]"
          }

###        location
Sub-section of type object, used to specify the location of the person
The schema snippet can be shown below:

       location:
        {
            "type": "object",
            "format": "location",
            "description": "e.g. Germany - [Location for this activity]",
            "properties": {
              "lat": {
                "type": "number"
              },
              "long": {
                "type": "number"
              }
            }
          }

###        url
Sub-section of type string, used to specify the url of the person
The schema snippet can be shown below:

       url:
        {
            "type": "string",
            "description": "e.g. http://xyz.example.com - [Related link to the company website]",
            "format": "uri"
          }

###        startDate
Sub-section of type string, used to specify the startDate of the person
The schema snippet can be shown below:

       startDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-28 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2018-12-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        summary
Sub-section of type string, used to specify the summary of the person
The schema snippet can be shown below:

       summary:
        {
            "type": "string",
            "description": "Give an overview of your responsibilities at the company"
          }

###        highlights
Sub-section of type array, used to specify the highlights of the person
The schema snippet can be shown below:

       highlights:
        {
            "type": "array",
            "description": "Specify multiple accomplishments",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. Worked with mobile team at Twitter to develop remote debugging tools for mobile browsers "
            }
          }


# Education
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalItems
##   items
###        institution
Sub-section of type string, used to specify the institution of the person
The schema snippet can be shown below:

       institution:
        {
            "type": "string",
            "description": "e.g. XYZ Institute of Technology - [Add institute name]"
          }

###        area
Sub-section of type string, used to specify the area of the person
The schema snippet can be shown below:

       area:
        {
            "type": "string",
            "description": "e.g. Engineering"
          }

###        studyType
Sub-section of type string, used to specify the studyType of the person
The schema snippet can be shown below:

       studyType:
        {
            "type": "string",
            "description": "e.g. Bachelor"
          }

###        startDate
Sub-section of type string, used to specify the startDate of the person
The schema snippet can be shown below:

       startDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2013-06-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        score
Sub-section of type object, used to specify the score of the person
The schema snippet can be shown below:

       score:
        {
            "type": "object",
            "additionalProperties": true,
            "properties": {
              "scoretype": {
                "type": "string",
                "description": "eg. GPA or PERCENTAGE - [Add score type]"
              },
              "scorevalue": {
                "type": "string",
                "description": "eg. 3.4/4.0 - [Add obtained score/total score]"
              }
            }
          }

###        courses
Sub-section of type array, used to specify the courses of the person
The schema snippet can be shown below:

       courses:
        {
            "type": "array",
            "description": "List notable courses/subjects",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. CS302 - Introduction to Algorithms - [Add course name]"
            }
          }

###        honors
Sub-section of type array, used to specify the honors of the person
The schema snippet can be shown below:

       honors:
        {
            "type": "array",
            "description": "List education honours",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. Magna Cum Laude"
            }
          }


# Volunteer
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalItems
##   items
###        organization
Sub-section of type string, used to specify the organization of the person
The schema snippet can be shown below:

       organization:
        {
            "type": "string",
            "description": "e.g. Xyz "
          }

###        position
Sub-section of type string, used to specify the position of the person
The schema snippet can be shown below:

       position:
        {
            "type": "string",
            "description": "e.g. Open Source Contributor [Contribution type]"
          }

###        location
Sub-section of type object, used to specify the location of the person
The schema snippet can be shown below:

       location:
        {
            "type": "object",
            "format": "location",
            "description": "e.g. Germany - [Location for this activity]",
            "properties": {
              "lat": {
                "type": "number"
              },
              "long": {
                "type": "number"
              }
            }
          }

###        url
Sub-section of type string, used to specify the url of the person
The schema snippet can be shown below:

       url:
        {
            "type": "string",
            "description": "e.g. http://xyz.example.com - [Related link to support volunteer experience]",
            "format": "uri"
          }

###        startDate
Sub-section of type string, used to specify the startDate of the person
The schema snippet can be shown below:

       startDate:
        {
            "type": "string",
            "description": "e.g. 2014-06-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ",
            "format": "date"
          }

###        summary
Sub-section of type string, used to specify the summary of the person
The schema snippet can be shown below:

       summary:
        {
            "type": "string",
            "description": "Give an overview of your responsibilities at the company"
          }

###        highlights
Sub-section of type array, used to specify the highlights of the person
The schema snippet can be shown below:

       highlights:
        {
            "type": "array",
            "description": "Specify accomplishments and achievements",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g Invited as a speaker in Xyzcon'17"
            }
          }


# Publications
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. Deep learning and Artificial Intelligence"
          }

###        publisher
Sub-section of type string, used to specify the publisher of the person
The schema snippet can be shown below:

       publisher:
        {
            "type": "string",
            "description": "e.g. XYZ, Computer Magazine"
          }

###        releaseDate
Sub-section of type string, used to specify the releaseDate of the person
The schema snippet can be shown below:

       releaseDate:
        {
            "type": "string",
            "description": "e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard]"
          }

###        resources
Sub-section of type array, used to specify the resources of the person
The schema snippet can be shown below:

       resources:
        {
            "type": "array",
            "description": "Specify multiple resources with label",
            "additionalItems": false,
            "url": {
              "type": "string",
              "format": "uri",
              "description": "e.g. http://www.example.com/my-example-slides/"
            },
            "label": {
              "type": "string",
              "description": "e.g Slides"
            }
          }

###        url
Sub-section of type string, used to specify the url of the person
The schema snippet can be shown below:

       url:
        {
            "type": "string",
            "description": "e.g. http://www.computer.org.example.com/csdl/mags/co/2015/10/rx069-abs.html"
          }

###        summary
Sub-section of type string, used to specify the summary of the person
The schema snippet can be shown below:

       summary:
        {
            "type": "string",
            "description": "e.g. Discussion of the advent of deep learning and artificial intelligence - Short summary of publication"
          }


# Legal
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. XYZ's patent on LZW compression, a fundamental part of the widely used GIF graphics format - [Add document name]"
          }

###        legalType
Sub-section of type string, used to specify the legalType of the person
The schema snippet can be shown below:

       legalType:
        {
            "type": "string",
            "description": "e.g. Patent, Trademark, Copyright - Give the type of this document"
          }

###        description
Sub-section of type string, used to specify the description of the person
The schema snippet can be shown below:

       description:
        {
            "type": "string",
            "description": "Give a brief description about this document"
          }

###        applicationDate
Sub-section of type string, used to specify the applicationDate of the person
The schema snippet can be shown below:

       applicationDate:
        {
            "type": "string",
            "description": "e.g. 2015-08-01 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        grantDate
Sub-section of type string, used to specify the grantDate of the person
The schema snippet can be shown below:

       grantDate:
        {
            "type": "string",
            "description": "e.g. 2016-09-01 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2020-09-03 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        resources
Sub-section of type array, used to specify the resources of the person
The schema snippet can be shown below:

       resources:
        {
            "type": "array",
            "description": "Specify multiple resources with label",
            "additionalItems": false,
            "url": {
              "type": "string",
              "format": "uri",
              "description": "e.g. http://www.example.com/my-patent-slides/"
            },
            "label": {
              "type": "string",
              "description": "e.g Slides"
            }
          }

###        idNumber
Sub-section of type string, used to specify the idNumber of the person
The schema snippet can be shown below:

       idNumber:
        {
            "type": "string",
            "description": "e.g. JP2004369746A - [Add the application number or Id Number]  "
          }


# Skills
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. Web Development"
          }

###        keyword
Sub-section of type array, used to specify the keyword of the person
The schema snippet can be shown below:

       keyword:
        {
            "type": "array",
            "description": "List some keywords pertaining to the skill",
            "additionalItems": false,
            "items": {
              "type": "object",
              "additionalProperties": true,
              "properties": {
                "name": {
                  "type": "string",
                  "description": "e.g. HTML - [Add the skill name]"
                },
                "score": {
                  "type": "number",
                  "description": "e.g. 20 - [Score for the skill name]"
                }
              }
            }
          }


# Awards
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        title
Sub-section of type string, used to specify the title of the person
The schema snippet can be shown below:

       title:
        {
            "type": "string",
            "description": "e.g. Awarded Software Process Achievement Award "
          }

###        date
Sub-section of type string, used to specify the date of the person
The schema snippet can be shown below:

       date:
        {
            "type": "string",
            "description": "e.g. 2016-06-12 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        awarder
Sub-section of type string, used to specify the awarder of the person
The schema snippet can be shown below:

       awarder:
        {
            "type": "string",
            "description": "e.g.  IEEE"
          }

###        summary
Sub-section of type string, used to specify the summary of the person
The schema snippet can be shown below:

       summary:
        {
            "type": "string",
            "description": "e.g. Received for my work in Deep learning and AI"
          }


# Projects
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. File Transfer application - [Name of the project]"
          }

###        location
Sub-section of type object, used to specify the location of the person
The schema snippet can be shown below:

       location:
        {
            "type": "object",
            "format": "location",
            "description": "e.g France - [Location for this activity]",
            "properties": {
              "lat": {
                "type": "number"
              },
              "long": {
                "type": "number"
              }
            }
          }

###        description
Sub-section of type string, used to specify the description of the person
The schema snippet can be shown below:

       description:
        {
            "type": "string",
            "description": "e.g. Developed a client and server based application - [Short summary of project]"
          }

###        highlights
Sub-section of type array, used to specify the highlights of the person
The schema snippet can be shown below:

       highlights:
        {
            "type": "array",
            "description": "Specify multiple features",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. used Java AWT and Swing for client side userinterface"
            }
          }

###        keywords
Sub-section of type array, used to specify the keywords of the person
The schema snippet can be shown below:

       keywords:
        {
            "type": "array",
            "description": "Specify special elements involved",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. Java"
            }
          }

###        startDate
Sub-section of type string, used to specify the startDate of the person
The schema snippet can be shown below:

       startDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard] ",
            "format": "date"
          }

###        resources
Sub-section of type array, used to specify the resources of the person
The schema snippet can be shown below:

       resources:
        {
            "type": "array",
            "description": "Specify multiple resources with label",
            "additionalItems": false,
            "url": {
              "type": "string",
              "format": "uri",
              "description": "e.g. http://www.example.com/my-awesome-slides/"
            },
            "label": {
              "type": "string",
              "description": "e.g slides"
            }
          }

###        url
Sub-section of type string, used to specify the url of the person
The schema snippet can be shown below:

       url:
        {
            "type": "string",
            "format": "uri",
            "description": "e.g. http://www.example.org/csdl/mags/co/1996/10/rx069-abs.html"
          }

###        roles
Sub-section of type array, used to specify the roles of the person
The schema snippet can be shown below:

       roles:
        {
            "type": "array",
            "description": "Specify your role on this project or in company",
            "additionalItems": false,
            "items": {
              "type": "string",
              "description": "e.g. Team Lead, Speaker, Writer"
            }
          }

###        entity
Sub-section of type string, used to specify the entity of the person
The schema snippet can be shown below:

       entity:
        {
            "type": "string",
            "description": "e.g. 'greenpeace', 'corporationXYZ' - [Relevant company/entity affiliations]"
          }

###        type
Sub-section of type string, used to specify the type of the person
The schema snippet can be shown below:

       type:
        {
            "type": "string",
            "description": " e.g. 'volunteering', 'presentation', 'talk', 'application', 'conference'"
          }


# Certificates
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalItems
##   items
###        code
Sub-section of type string, used to specify the code of the person
The schema snippet can be shown below:

       code:
        {
            "type": "string",
            "description": "e.g. 1Z0-062"
          }

###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. XYZ Certified Application Specialist (MCAS) - [Add the certificate name]"
          }

###        website
Sub-section of type string, used to specify the website of the person
The schema snippet can be shown below:

       website:
        {
            "type": "string",
            "description": "Link to issuing authority's description of the certificate",
            "format": "uri"
          }

###        verification
Sub-section of type string, used to specify the verification of the person
The schema snippet can be shown below:

       verification:
        {
            "type": "string",
            "description": "External candidate verification URL",
            "format": "uri"
          }

###        grantDate
Sub-section of type string, used to specify the grantDate of the person
The schema snippet can be shown below:

       grantDate:
        {
            "type": "string",
            "description": "e.g. 2017-06-29 - [resume.json uses the ISO 8601 date standard]",
            "format": "date"
          }

###        score
Sub-section of type string, used to specify the score of the person
The schema snippet can be shown below:

       score:
        {
            "type": "string",
            "description": "Exam result (PASS/FAIL, 100%, 100)",
            "format": "date"
          }

###        endDate
Sub-section of type string, used to specify the endDate of the person
The schema snippet can be shown below:

       endDate:
        {
            "type": "string",
            "description": "e.g. 2020-01-20",
            "format": "date"
          }

###        doesNotExpire
Sub-section of type boolean, used to specify the doesNotExpire of the person
The schema snippet can be shown below:

       doesNotExpire:
        {
            "type": "boolean",
            "format": "checkbox"
          }


# References
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. Stephan Mark"
          }

###        company
Sub-section of type string, used to specify the company of the person
The schema snippet can be shown below:

       company:
        {
            "type": "string",
            "description": "e.g. Xyz"
          }

###        position
Sub-section of type string, used to specify the position of the person
The schema snippet can be shown below:

       position:
        {
            "type": "string",
            "description": "e.g. Senior Software Engineer"
          }

###        reference
Sub-section of type string, used to specify the reference of the person
The schema snippet can be shown below:

       reference:
        {
            "type": "string",
            "description": "e.g. Joe blogs was a great employee, who turned up to work at least once a week. He exceeded my expectations when it came to doing nothing."
          }


# Languages
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalItems
##   items
###        language
Sub-section of type string, used to specify the language of the person
The schema snippet can be shown below:

       language:
        {
            "type": "string",
            "description": "e.g. English, Spanish - [Name of language]"
          }

###        score
Sub-section of type number, used to specify the score of the person
The schema snippet can be shown below:

       score:
        {
            "type": "number",
            "description": "e.g. 20 - [Score for the language]"
          }


# Interests
This section is of array type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   additionalItems
##   items
###        name
Sub-section of type string, used to specify the name of the person
The schema snippet can be shown below:

       name:
        {
            "type": "string",
            "description": "e.g. Philosophy"
          }


# Meta
This section is of object type that tells about the basic information of the user and consists of the following sub-sections:
##   type
##   description
##   additionalProperties
##   properties
###        canonical
Sub-section of type string, used to specify the canonical of the person
The schema snippet can be shown below:

       canonical:
        {
          "type": "string",
          "description": "URL (as per RFC 3986) to latest version of this document"
        }

###        version
Sub-section of type string, used to specify the version of the person
The schema snippet can be shown below:

       version:
        {
          "type": "string",
          "description": "e.g. v1.0.0 - A version field which follows semver"
        }

###        lastModified
Sub-section of type string, used to specify the lastModified of the person
The schema snippet can be shown below:

       lastModified:
        {
          "type": "string",
          "description": "Using ISO 8601 with YYYY-MM-DDThh:mm:ss"
        }

