# Task Tracker

## Description
This is a task tracker project. The goal is to build a service that uses
a markdown based todo-list task-organization and tracks the task details like
start time, task status changes, tags, estimates, revised estimates, life proj
etc. and maintains a databse with these details across devices.
The goal of the project is to come up with a hierarchical data model that allows
data collection and visualization at various levels.

Task tracker is the first among a variety of productivity/accountability tracker
that will help individuals and self organized teams a deeper look into their
productivity graphs and address blocking issues head on.

The task tracker is essentially a collection of services that monitor "registered" 
todo files and collect, format, and post data to the level data store at current
level of hierarchy. Also includes data aggregation/purge services that are specific
to todo tasks.

It will include a service that takes the current data and visualizes it as 
bunch of scrum cards per project.

[Possible approach 1]
The Task tracker converts and stores the data using the logstash format in order
to be able to use the existing ELK stack for data visualization and querying.
In this case it will become an elk extension of some sort.. which won't use
most of the elk stack capabilities, in fact maybe rewrite some of them.

Advantage of using EKL stack is that it is already widely accepted, and maintained
and it should be relatively simple to onboard people

[Possible approach 2]
Since kibana is highly coupled with logstash and elasticsearch, and I want to 
use my own data model and data aggregation, it might not be the best way to go.
In that case I need to build either a kibana connector that can proxy this part
or I need some other visualization tool.. which should be open source. Grafana
may be a decent choice but search capability would need to be developed separately.

[Production case]
build a basic custom data aggregation and visuzlisation [mvp]
support plugability of various types of data aggregators/visualizers in addition
to the basic one.

[Goal]
The aim is to use this as a vim or a VScode extension across multiple devices.
This also includes tasks that I write at the end of my code files maybe


[challenges]
- reading change status of a file as a service
- posting to database -> different devices, connectivity issues, authentication etc
- hierarchical data forat
- building dashboards on top of data/tasks
