I have created /secretpage route which needs authentication to send some data.
If you are not logged in than this page will send an login request.

The token will expire automatically in 15mins and user has to login again after this
And token will expire only after this time (i will add redis to resolve this two problems)
