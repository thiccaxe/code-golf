package routes

import (
	"net/http"

	"github.com/code-golf/code-golf/hole"
	"github.com/code-golf/code-golf/lang"
	"github.com/code-golf/code-golf/pager"
	"github.com/code-golf/code-golf/session"
)

// RankingsMedals serves GET /rankings/medals/{hole}/{lang}
func RankingsMedals(w http.ResponseWriter, r *http.Request) {
	type row struct {
		Country, Login             string
		Rank, Gold, Silver, Bronze int
	}

	data := struct {
		Holes []hole.Hole
		Langs []lang.Lang
		Pager *pager.Pager
		Rows  []row
	}{
		Holes: hole.List,
		Langs: lang.List,
		Pager: pager.New(r),
	}

	rows, err := session.Database(r).Query(
		`SELECT RANK() OVER(ORDER BY gold DESC, silver DESC, bronze DESC),
		        COALESCE(CASE WHEN show_country THEN country END, ''),
		        login,
		        gold,
		        silver,
		        bronze,
		        COUNT(*) OVER()
		   FROM medals
		   JOIN users ON id = user_id
		  LIMIT $1 OFFSET $2`,
		pager.PerPage,
		data.Pager.Offset,
	)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var r row

		if err := rows.Scan(
			&r.Rank,
			&r.Country,
			&r.Login,
			&r.Gold,
			&r.Silver,
			&r.Bronze,
			&data.Pager.Total,
		); err != nil {
			panic(err)
		}

		data.Rows = append(data.Rows, r)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	if data.Pager.Calculate() {
		NotFound(w, r)
		return
	}

	render(w, r, "rankings/medals", "Rankings: Medals", data)
}
