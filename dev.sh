air 2>&1 | awk '{print "air:      " $0}' &
templ generate -watch ./ 2>&1 | awk '{print "templ:    " $0}' &
tailwindcss -w ./ -i ./views/input.css -o ./public/output.css 2>&1 | awk '{print "tailwind: " $0}'
