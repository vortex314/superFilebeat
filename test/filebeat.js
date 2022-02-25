console.log("Javascript engine loaded ");
process = function (event) {
    console.log(">>>> " + JSON.stringify(event.Fields))
    console.log(event.Fields.message)
    event.Fields.metrics = {}
    if (!isNaN(parseFloat(event.Fields.upstream_response_time))) event.Fields.metrics.http_response_time = parseFloat(event.Fields.upstream_response_time)
    if (!isNaN(parseFloat(event.Fields.body_bytes_sent))) event.Fields.metrics.http_bytes_sent = parseFloat(event.Fields.body_bytes_sent)
    event.Fields.attributes.http_response_code = event.Fields.errorCode
    //event.Fields.metrics.age=60
    //console.log("")
    event.Fields.message = "" // dropped source message, kafka complained about size > 512
    //console.log("<<<<<< " +JSON.stringify(event.Fields))
    //console.log("")
    //console.log(" BEFORE "+JSON.stringify(event.Fields.upstream_response_time))
    //console.log(" METRICS "+JSON.stringify(event.Fields.metrics))
    return event
}