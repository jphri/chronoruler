# ChronoRuler

Não é uma agenda, é uma régua que ajuda você a alocar o seu tempo.

## Racionalização
### A Régua

Comecemos a pensar em um dia como um banco de horas contendo 24 horas de
reserva que podemos usar. Dito isso, podemos começar a adicionar nossas
necessidades e projetos que devemos fazer.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>8h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>8h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>8h</td>
        </tr>
    </tfoot>
</table>

Porém, nós podermos analisar algumas tarefas de forma a fazer uma proporção.
Por exemplo, dormir toma 1/3 do dia em um ano. Se nós escalonarmos todos os
valores para 1 ano, nós temos:

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>2920h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>2920h</td>
        </tr>
    </tfoot>
</table>

Vamos considerar que nesse ponto queremos fazer um projeto X que sabemos que
podemos terminar em mais ou menos 48 horas. Se nós considerarmos que esse é um
projeto que queremos fazer dentro desse ano, nós podemos adicionar só mais um
projeto de 48 horas nesse sistema de horas por ano.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Projeto X</td>
            <td>48h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>2872h</td>
        </tr>
    </tfoot>
</table>

E depois escalonar esse sistema de horas para os 24 horas por dia.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>8h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>8h</td>
        </tr>
        <tr>
            <td>Projeto X</td>
            <td>0.13h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>7.87h</td>
        </tr>
    </tfoot>
</table>

O que nos ajuda a perceber que, se quisermos terminar o projeto X em 1 ano, nós
podemos trabalhar 0.13 horas por dia (≃ 8 minutos por dia).

### Prioridades e tempo de termino

Nós podemos representar prioridades em quantas vezes nós queremos terminar esse
projeto dentro de um período de tempo. Voltemos para a tabela de horas por ano.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Projeto X</td>
            <td>48h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>2872h</td>
        </tr>
    </tfoot>
</table>

Vamos dizer que que queremos terminar esse projeto em 6 meses. Ao invés de
interpretar que queremos fazer esse projeto em 6 meses, podemos interpretar que
queremos fazer esse projeto 2 vezes em 1 ano, e isso custa o dobro do tempo em
1 ano.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>2920h</td>
        </tr>
        <tr>
            <td>Projeto X</td>
            <td>96h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>2824h</td>
        </tr>
    </tfoot>
</table>

E se escalonarmos de volta para 24 horas por dia.

<table>
    <tbody>
        <tr>
            <th>Ação</th>
            <th>Horas</th>
        </tr>
        <tr>
            <td>Dormir</td>
            <td>8h</td>
        </tr>
        <tr>
            <td>Trabalhar</td>
            <td>8h</td>
        </tr>
        <tr>
            <td>Projeto X</td>
            <td>0.26h</td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <th>Horario restante</th>
            <td>7.74h</td>
        </tr>
    </tfoot>
</table>

Isso significa que nós devemos trabalhar 0.26h por dia pra terminarmos esse
mesmo projeto em 6 meses (≃ 16 minutos).

## Uso

```
./chronoruler -config=<arquivo.json> [-mode=<mode>]
mode:
    - hpd = horas por dia (hours per day)
    - hpy = horas por ano (hours per year) (padrão)
    - mpd = minutos por dia (minutes per day)
```

