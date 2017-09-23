#Overview
The schema.json file consists of a set of sections,sub-sections with related properties.
The current version of schema.json consists of the following sub-sections:

    * core
    * work
    * education
    * courses
    * volunteer
    * publications
    * legal
    * skills
    * awards
    * projects
    * certificates
    * references
    * languages
    * interests
    * meta

# core

This section is of object type that tells about the basic information of the user and consists of the following sub-sections:

##### name  
    Sub-section of type string, used to specify the name of the person.
##### title
    Sub-section of type string, used to specify the current title of the person.
##### image
    Sub-section of type string, used to specify the url of an image (as per RFC 3986)in JPEG or PNG format.
##### email
    Sub-section of type string, used to specify the url of an image (as per RFC 3986)in JPEG or PNG format.
##### phone
    Sub-section of type string, used to specify the contact number of the user.
##### url
    Sub-section of type string, used to specify the contact number of the user.
##### summary
    Sub-section of type string used to write a short 2-3 sentence biography about yourself
##### currentLocation
    Sub-section of type object used to select location.
    It has two properties lat and long:
        "lat": {
              "type": "number"
            },
            "long": {
              "type": "number"
            }
    which will store the latitude and longitude numerical values of the current location. 
##### permanentLocation
    Sub-section of type object used to select location.
    It has two properties lat and long:
        "lat": {
              "type": "number"
            },
            "long": {
              "type": "number"
            }
    which will store the latitude and longitude numerical values of the permanent location. 
