# Itinerary app

## Use case note 
The *itinerary app* reads in two files, one with airport data under the assets directory and an administrative text document that contains the data to be parsed and saved as the end customer document. 

The input and output is stored in the  *test* directory and the airport data base is located  in the *assets* directory.  

 **The app is not compiled as an executable, you must have GoLang installed to run the app**  

## Execution instructions

- 1) To execute the program run $go run . -arg1 -arg2 -arg3 in while in the directory the program installed is installed in.
- 2) Use the run_itinerary.sh script (this will run the program with preset arguments and remove the output file ) 

## The arguments

* The input (itinerary admin file) will read in a a file as per the specified file path that includes the file name ant type must be passed in as **-arg1**. 

* The out put of the file will stored to file path and name as has been that includes the file name and type. must be passed in as **-arg2** 

*  airport data base will be read in from  **-arg3**   

## Operations

- will convert #AAA iata code to the corresponding airport name 
- will convert ##AAAA icao to the corresponding airport name
- will convert *AAA iata  code to  the corresponding municipality name  
- will convert *#AAAA icao code to  the corresponding municipality name
- will convert D(ISO 8601)  time to the date in the format of DD-Mmm-YYYY
- will convert T12(ISO 8601)  time to the the AM/PM format followed by the timezone in parentheses
- will convert T24(ISO 8601)  time to the the 24 format followed by the timezone in parentheses  
- will remove all but one empty lines
- will remove the following white spaces if left in to the text "\v" "\f" "\r" and substitute it for "\n"  
 
## Restrictions

-The Administrator file will be read in ***AS-IS***  meaning only the conversion will be made as stated in the operations
-The Time and Date will always need to be followed by a space, else the time will not be registered as a time 
-If there is any error that occurs the output file will not be created
-If an error occurs the error will be printed to the console 

## Contact 

Have questions, concerns or issues related to the itinerary program pleas contact Hannes at  hannes.study@gmail.com

bugs should be report to hannes.study@gmail.com