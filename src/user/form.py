from wtforms import StringField

from common.form import PagenationForm


class QueryUserForm(PagenationForm):
    name = StringField()
    email = StringField()
