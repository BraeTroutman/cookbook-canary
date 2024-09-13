---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

layout: home
---

{% for recipe in site.recipes %}
[{{ recipe.title }}]({{ recipe.url | prepend: site.baseurl }})
{% endfor %}

