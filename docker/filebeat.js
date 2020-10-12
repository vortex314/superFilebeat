console.log("Javascript engine loaded "); 
process = function(event){
   console.log("DATES  ---> "+JSON.stringify(event.Timestamp)+" "+JSON.stringify(event.Private.Timestamp)+" "+JSON.stringify(event.Fields.timestamp))
   console.log("FIELDS ---> "+JSON.stringify(event.Fields))
   event.Fields.metrics={}
 //   event.Fields.metrics.age=60
 //   console.log("")
//      event.Fields.message="" // dropped source message, kafka complained about size > 512
//    console.log("<<<<<< " +JSON.stringify(event.Fields))
 //   console.log("")
 //   console.log(" BEFORE "+JSON.stringify(event.Fields.upstream_response_time))
 //   console.log(" METRICS "+JSON.stringify(event.Fields.metrics))
	return event
	}
