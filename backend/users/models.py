from django.db import models
from django.utils import timezone
from django.utils.translation import gettext_lazy as _
from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin, BaseUserManager

from .utils import generate_random_password


# Create your models here.
class CustomAccountManager(BaseUserManager):
    
    def create_user(self, email: str, user_name: str, password: str = None, **other_fields) -> 'User':

        if not email:
            raise ValueError('Users must provide an email address')

        email = self.normalize_email(email)
        user = self.model(
            email=email,
            user_name=user_name,
            **other_fields
        )
        if password is None:
            password = generate_random_password()

        user.set_password(password)
        user.save()
        return user
    
    def create_superuser(self, email: str, user_name: str, password: str = None, **other_fields) -> 'User':
        
        other_fields.setdefault('is_staff', True)
        other_fields.setdefault('is_superuser', True)
        other_fields.setdefault('is_active', True)

        if other_fields.get('is_staff') is not True:
            raise ValueError('Superuser must have is_staff=True.')
        
        if other_fields.get('is_superuser') is not True:
            raise ValueError('Superuser must have is_superuser=True.')
        
        return self.create_user(
            email=email,
            user_name=user_name,
            password=password,
            **other_fields)
    


class User(AbstractBaseUser, PermissionsMixin):

    email = models.EmailField(_('Email'), unique=True)
    user_name = models.CharField(max_length=15, unique=True)
    is_staff = models.BooleanField(default=False)
    is_active = models.BooleanField(default=False)
    date_joined = models.DateTimeField(default=timezone.now)
    

    class Meta:
        db_table = 'users'

    objects = CustomAccountManager()

    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['user_name']

    def __str__(self) -> str:
        return self.user_name