from django.urls import path

from . import views

app_name = "teamjj"
urlpatterns = [
 	# ex: /teamjj/
    path('', views.index, name='index'),
    # ex: /teamjj/user_name/
    path('<user_name>/', views.user_detail, name='user_detail'),
    # ex: /teamjj/projects/1/
    path('projects/<int:id>/', views.project_detail, name='project_detail'),
  
]