def main():
    facebook_auth = "https://graph.facebook.com/oauth/access_token?client_id=571605572000175&client_secret=c3cc928991d69e784f69a907000cc909&grant_type=client_credentials"
    facebook_access_token = "https://graph.facebook.com/v2.8/218718188140085/members?access_token=EAACEdEose0cBAMF99BFdceBHS5UcaEPIDbhKT1nbE0DaulfIS8ZCJTHrkD6IxkfKZCObCHzUsHcdhBISYuv7L4oiylZCmNztRy6s9KBgmZBYbJxZBBk1Csjl8PXFnZA3rt5WwZB7dZCB9M04a6FI7XwCFXfNV4ZB76y6SjltVc4AP4gZDZD&pretty=0&fields=id&limit=1000&after=MzQxNDM1MTE2MjQ0Njk4"
    return facebook_auth, facebook_access_token
