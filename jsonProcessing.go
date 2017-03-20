package main

import (
  "encoding/json"
  "strings"
  "github.com/golang/glog"
)

func toJson(template interface{}) string {
  bytes, err := json.Marshal(template)
  if err != nil {
    glog.Fatal(err.Error())
  }

  return string(bytes)
}

func isJson(s string) (map[string]interface{}, bool) {
  var js map[string]interface{}

  err := json.Unmarshal([]byte(s), &js)

  return js, err == nil
}

func getPathFromJson(pathJson []byte) (paths []string) {
  pathString := string(pathJson)

  glog.V(1).Info(pathString) 

  return paths
}

func getPathAndVerbFromJson(tag string, operationID string) (path string, verb string) {
  glog.V(1).Info(operationID)
  context, verb := getContextAndVerb(removeQuotes(string(operationID)), removeQuotes(string(tag)))
  path = "/" + removeQuotes(string(tag)) + context
  return path, verb
}

func getContextAndVerb(operationID string, tag string) (context string, verb string) {

  switch true {
    case strings.HasPrefix(operationID, "add"):
      context = strings.TrimPrefix(operationID, "add")
      verb = "post"
    case strings.HasPrefix(operationID, "post"):
      context = strings.TrimPrefix(operationID, "post")
      verb = "post"
    case strings.HasPrefix(operationID, "upload"):
      context = strings.TrimPrefix(operationID, "upload")
      verb = "post"
    case strings.HasPrefix(operationID, "update"):
      context = strings.TrimPrefix(operationID, "update")
      verb = "put"
    case strings.HasPrefix(operationID, "get"):
      context = strings.TrimPrefix(operationID, "get")
      verb = "get"
    case strings.HasPrefix(operationID, "find"):
      context = operationID
      verb = "get"
    case strings.HasPrefix(operationID, "delete"):
      context = strings.TrimPrefix(operationID, "delete")
      verb = "delete"
    default:
      context = operationID
      verb = "get"
    }
    glog.V(1).Info("Verb is : " + verb)
    glog.V(1).Info("Tag is : " + tag)

    //add support for plural tag
    contextSplit := strings.Split(context, strings.Title(tag))
    if strings.Contains(context, strings.Title(tag)) && strings.HasPrefix(contextSplit[1], "s") {
      context = strings.Replace(context, strings.Title(tag) + "s", "", 1)
    } else {
      context = strings.Replace(context, strings.Title(tag), "", 1)
    }
    context = strings.Replace(context, " ", "_", -1)
    glog.V(1).Info("Context is : " + context)
    if context != "" {
      context = "/" + context
    }

  return context, verb
}

/**func getSwaggerTemplate() SwaggerTemplate {
  raw, err := ioutil.ReadFile("./Swagger_UI_Poc.json")
  if err != nil {
    glog.Fatal(err.Error())
  }

  var c SwaggerTemplate
  json.Unmarshal(raw, &c)
  return c
}**/

func removeQuotes(quotedString string) (unquotedString string) {
  unquotedString = strings.TrimPrefix(quotedString, "\"")
  return strings.TrimSuffix(unquotedString, "\"")
}

func removeWhiteSpace(stringWithSpaces string) (stringWithoutSpaces string) {
  stringWithoutSpaces = strings.Replace(stringWithSpaces, "\n", "", -1)
  stringWithoutSpaces = strings.Replace(stringWithoutSpaces, "\r", "", -1)
  stringWithoutSpaces = strings.Replace(stringWithoutSpaces, "\t", "", -1)
  return stringWithoutSpaces
}
