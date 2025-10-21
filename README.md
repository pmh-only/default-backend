# default-backend
default backend service for pmh.codes homelab.

## wdym?
this implements [default backend](https://kubernetes.github.io/ingress-nginx/user-guide/default-backend/) and [custom errors](https://kubernetes.github.io/ingress-nginx/user-guide/custom-errors/) service specs for ingress nginx controller

## i wanna try
ok here: https://goofy-ahh.endpoint-bruhhhhhhh.pmh.codes (this causes 404 error, the intended behavior)

you need to bypass HSTS preload, since my certificate doesn't cover that url.\
[simply type `thisisunsafe` in ssl error screen if you use chrome](https://www.reddit.com/r/webdev/comments/kzgozy/til_you_can_type_thisisunsafe_on_chrome_ssl_error/)

or use less goofy url: https://doesntexist.mc.pmh.codes

## no, that's not what i mean. i wanna run it
why?

uhh okay. then just run `ghcr.io/pmh-only/default:latest` image. \
and follow [ingress nginx controller's official document](https://kubernetes.github.io/ingress-nginx/examples/customization/custom-errors/#ingress-controller-configuration)

oh. i almost forgot to tell you. server is listening on `3000` port.

## huh, so this thing just exist for displaying fancier 404 page than nginx's default one
yes. bingo.

actually, it handles 5xx error also when something crashed or timeout.

## licenses?
&copy; Minhyeok Park. 2025. 0BSD Licensed.
