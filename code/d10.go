${?#{_d10}}$
${#roll:= [1d10]}$
${#total:= roll + modificateur}$

${#result:= total <= 1 ? "Non et..." : total <= 3 ? "Non" : total <= 5 ? "Non, mais..." : total <= 7 ? "Oui, mais..." : total <= 9 ? "Oui" : "Oui et..."}$

${#class:= total <= 1 ? "non-et" : total <= 3 ? "non" : total <= 5 ? "non-mais" : total <= 7 ? "oui-mais" : total <= 9 ? "oui" : "oui-et"}$

<div class="omenxx-result-detail"><i class="fa-regular fa-dice-d10"></i>${!roll}$ ${!modificateur >=0 ? "+" : "-"}$ ${!abs(modificateur)}$ = <strong>${!total}$</strong></div>

<div class="omenxx-result omenxx-result-${!class}$">
  ${!result}$
</div>
