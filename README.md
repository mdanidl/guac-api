# Guacamole Golang Frontend API

This package is built for (Apache Guacamole)[http://guacamole.apache.org].

I created it to access the built-in client application's frontend API in order to programatically add Users, Connections and Permissions for those connections.
It's done by reverse engineering the AJAX calls from the webUI.

The Goal is to use this library in a terraform provider that creates these resources when needed.

In it's current state the style of the code is very raw, and a **Work In Progress**
Also a caveat that i am by no means a software developer, and i'm doing this partly to get familiar with Golang.

PRs are most welcome.