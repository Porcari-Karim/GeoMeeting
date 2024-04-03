from django.contrib import admin
from .models import User
from django.contrib.auth.admin import UserAdmin
from django.forms import TextInput, Textarea

# Register your models here.
class UserAdminConfig(UserAdmin):
    ordering = ('-date_joined',)
    search_fields = ('email', 'user_name')
    list_filter = ('is_active', 'is_staff', 'email', 'user_name')
    list_display = ('email', 'user_name', 'is_active', 'is_staff')

    fieldsets = (
        (None, {
            'fields' : ('email', 'user_name')
        }),
        ('Permissions', {
            'fields' : ('is_active', 'is_staff', 'is_superuser', 'groups', 'user_permissions')
        })
    )

    add_fieldsets =( 
        (None, {
            'classes' : ('wide',),
            'fields' : ('email', 'user_name', 'password1', 'password2', 'is_active', 'is_staff')
        }),
        )


admin.site.register(User, UserAdminConfig)