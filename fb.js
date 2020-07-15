console.log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>> Javascript engine loaded "); 
process = function(event){
	fields = event.Fields
	fields.message = fields.syslog_message
	console.log(" JSON BEFORE  : " +JSON.stringify(fields))
	fields.attributes={}
	fields.attributes.text = "just some Javascript code"
	fields.metrics={}
	fields.javascript="running in GO!"
//	var d = new Date(fields.timestamp)
//	console.log(" date " + d)
	console.log(" JSON AFTER : " +JSON.stringify(event.Fields))
	return event
	}
