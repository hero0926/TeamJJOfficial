from django.db import models

# Create your models here.
class Skill(models.Model) :
	skill_name = models.CharField(max_length=100)
	skill_proficiency = models.IntegerField(default=0)
	skill_text = models.CharField(max_length=5000)

class User(models.Model) :
	user_name = models.CharField(max_length=100)
	user_type = models.CharField(max_length=100)
	user_age = models.IntegerField()
	user_phone = models.CharField(max_length=100)
	user_skills = models.ForeignKey(Skill, on_delete=models.CASCADE)

class Project(models.Model) :
    project_contributer = models.ForeignKey(User, on_delete=models.CASCADE)
    project_skills = models.ForeignKey(Skill, on_delete=models.CASCADE)
    project_name = models.CharField(max_length=200)
    project_text = models.CharField(max_length=5000)
    project_link = models.CharField(max_length=200) 
    pub_date = models.DateTimeField('date published')
