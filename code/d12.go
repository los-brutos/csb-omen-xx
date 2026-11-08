${?#{_d12}}$
${#roll:= [1d12]}$
${#total:= roll + modificateur}$

${#result:= total <= 2 ? "Non et..." : total <= 4 ? "Non" : total <= 6 ? "Non, mais..." : total <= 8 ? "Oui, mais..." : total <= 10 ? "Oui" : "Oui et..."}$

${#class:= total <= 2 ? "non-et" : total <= 4 ? "non" : total <= 6 ? "non-mais" : total <= 8 ? "oui-mais" : total <= 10 ? "oui" : "oui-et"}$

<div class="omenxx-result-detail"><i class="fa-regular fa-dice-d12"></i>${!roll}$ ${!modificateur >=0 ? "+" : "-"}$ ${!abs(modificateur)}$ = <strong>${!total}$</strong></div>

<div class="omenxx-result omenxx-result-${!class}$">
  ${!result}$
</div>
