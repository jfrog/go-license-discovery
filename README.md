# go-license-discovery

The license discovery is a library and set of tools
that can analyze license text to determine what type of license it contains.
It searches for license texts in a file and compares them to an archive of known licenses

This library encapsulate logic from 2 best licenses detection libraries :

-  License Classifier  : https://github.com/google/licenseclassifier
-  go-license-detector : https://github.com/src-d/go-license-detector

In Addition this library uses :
 - checksum based cache to have better performance detection of licenses which already been met
 - analyzed pom file to detect licenses which presented in the pom as a comment

Example :

copy licenses
