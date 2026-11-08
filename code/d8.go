${?#{_d8}}$
${#roll:= [1d8]}$

${#total:= roll + modificateur}$

${#result:= total <= 1 ? "Non et..." : total <= 3 ? "Non" : total <= 4 ? "Non, mais..." : total <= 5 ? "Oui, mais..." : total <= 7 ? "Oui" : "Oui et..."}$

${#class:= total <= 1 ? "non-et" : total <= 3 ? "non" : total <= 4 ? "non-mais" : total <= 5 ? "oui-mais" : total <= 7 ? "oui" : "oui-et"}$

<div class="omenxx-result-detail"><i class="fa-regular fa-dice-d8"></i>${!roll}$ ${!modificateur >=0 ? "+" : "-"}$ ${!abs(modificateur)}$ = <strong>${!total}$</strong></div>

<div class="omenxx-result omenxx-result-${!class}$">
  ${!result}$
</div>
