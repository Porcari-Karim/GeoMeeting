from django.http import HttpRequest, HttpResponse
from django.shortcuts import render, redirect
from django.contrib.auth import authenticate, login, logout
from django.contrib import messages
# Create your views here.
def login(request: HttpRequest) -> HttpResponse:
    return render(request, 'login.html')