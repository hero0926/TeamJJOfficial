from django.contrib import admin
from .models import Skill, User, Project

# Register your models here.

admin.site.register(Skill)
admin.site.register(User)
admin.site.register(Project)