/* example */
user example_dapr
db.createUser({user:"example",pwd:"123456",roles:[{role:"userAdmin",db:"example_eventstore"},{role:"dbAdmin",db:"example_eventstore"},{role:"readWrite",db:"example_eventstore"}]});

db.createCollection("dapr_app_logs")
db.createCollection("dapr_event_logs")

db.createCollection("system_message");
db.createCollection("test_aggregate");
db.createCollection("test_event");

use example_query
db.createUser({user:"example",pwd:"123456",roles:[{role:"userAdmin",db:"example_query"},{role:"dbAdmin",db:"example_query"},{role:"readWrite",db:"example_query"}]});
db.createCollection("inventory");
db.createCollection("sale_bill");
db.createCollection("sale_item");
db.createCollection("user");





