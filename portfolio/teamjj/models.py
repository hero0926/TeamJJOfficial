import datetime

from django.db import models
from django.utils import timezone

# Create your models here.
class Skill(models.Model) :


	skillProficiency = (

	    ('1', 'Freshman'),

	    ('2', 'Sophomore'),

	    ('3', 'Junior'),

	    ('4', 'Senior'),

	    ('5', 'Graduate'),

	)

	skill_profiency = models.CharField(max_length=1, choices=skillProficiency)
	skill_name = models.CharField(max_length=100)
	skill_text = models.CharField(max_length=5000)

	def __str__(self):
		return self.skill_name

class User(models.Model) :
	user_name = models.CharField(max_length=100)
	user_type = models.CharField(max_length=100)
	user_age = models.IntegerField()
	user_skills = models.ManyToManyField(Skill)
	user_phone = models.CharField(max_length=100)

	def __str__(self):
		return self.user_name

class Project(models.Model) :

	project_skills = models.ManyToManyField(Skill)
	project_contributer = models.ManyToManyField(User)

	project_name = models.CharField(max_length=200)
	project_text = models.CharField(max_length=5000)
	project_link = models.CharField(max_length=200) 
	pub_date = models.DateTimeField('date published')
	def __str__(self):
		return self.project_name

	def was_published_this_year(self):
		return self.pub_date >= timezone.now() - datetime.timedelta(days=365)
