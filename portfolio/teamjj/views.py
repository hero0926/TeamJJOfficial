from django.shortcuts import get_object_or_404, render
from django.http import HttpResponse
from django.template import loader


from .models import Project

# Create your views here.
def index(request):

	project_list = Project.objects.order_by("-pub_date")[:5]
	template = loader.get_template("teamjj/index.html")
	context = {
		"project_list" : project_list,
	}

	#return render(requeset, "teamjj/index.html", context)
	return HttpResponse(template.render(context, request))


def user_detail(request, user_name) :

    response = "You're looking at the user of %s."
    return HttpResponse(response % user_name)


def project_detail(request, id) :

	# project = get_object_or_404(Project, pk=id)
	try :
		project = Project.objects.get(pk=id)
	except Project.DoesNotExist :
		raise Http404("No Projects")
	return render(request, "teamjj/projectDetail.html", {"project" : project})


