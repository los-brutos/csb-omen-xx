${?#{_d6_fu_suv}}$
${#d1:= [1d6[black]]}$
${#d2:= [1d6[grey]]}$
${#d3:= [1d6[red]]}$

<!--
  résultats : 
    1 => non et
    2 => non
    3 => non mais
    4 => oui mais
    5 => oui
    6 => oui mais
--> 

${#total:= or(modificateur == -1, modificateur == 1) ? d1 + modificateur : modificateur == 2 ? max(d1, d2) : modificateur == 3 ? max(d1, d2, d3) : modificateur == -2 ? min(d1, d2) : modificateur == -3 ? min(d1, d2, d3) : d1}$

${#result:= total <= 1 ? "Non et..." : total <= 2 ? "Non" : total <= 3 ? "Non, mais..." : total <= 4 ? "Oui, mais..." : total <= 5 ? "Oui" : "Oui et..."}$

${#class:= total <= 1 ? "non-et" : total <= 2 ? "non" : total <= 3 ? "non-mais" : total <= 4 ? "oui-mais" : total <= 5 ? "oui" : "oui-et"}$

<div class="omenxx-result-detail"><i class="fa-regular fa-dice-d6"></i>${!d1}$ ${!abs(modificateur) >= 2 ? concat('<i class="fa-regular fa-dice-d6" style="color: grey;"></i>', d2) : ""}$ ${!abs(modificateur) == 3 ? concat('<i class="fa-regular fa-dice-d6" style="color: red;"></i>', d3) : ""}$ ${!modificateur == 1 ? concat("+ ",modificateur) : modificateur == -1 ? concat("- ",abs(modificateur)) : ""}$ = <strong>${!total}$</strong></div>

<div class="omenxx-result omenxx-result-${!class}$">
  ${!result}$
</div>
