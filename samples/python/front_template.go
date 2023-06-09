package pythonsamples

type FrontTemplate struct {}


func (front *FrontTemplate) IndexTemplate() string {
return`{% extends "layouts/front-layout.html" %}

{% block title %} Index {% endblock %}

{% block content %}

    <h2>MVC App with Flask</h2><br>

    <div class="alert ">
        <h6> Store your daily tasks</h6><hr>
        <p><strong>Features:</strong></p>
        <ul>
            <li>Show Tasks</li>
            <li>Add Tasks</li>
			<li>Edit Tasks</li>
			<li>Search Tasks</li>
            <li>Add User and Upload Image</li>
            <li>Tasks details</li>
        </ul>

        <p><strong>Authentication:</strong></p>
        <ul>
            <li>Register <a href="/register"> <b>Here</b> </a> </li>
            <li>Fill Register Form</li>
            <li>Enter with Username and Password <a href="/login"> <b>Here</b> </a></li>
        </ul>
    </div>
    {% endblock %}` 
}