package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/phoops/transitland-route-geometry-generator/internal/infrastructure/postgres"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

const (
	dbConnectionString = "host=localhost port=5432 dbname=gtfsdb user=transit password=transit sslmode=disable timezone=UTC"
)

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	testLogger = l.Sugar()
}

func TestClientRouteShapesCalculationNoRows(t *testing.T) {
	pgClient := sqlx.MustConnect("postgres", dbConnectionString)

	defer func() {
		err := pgClient.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := postgres.NewClient(
		testLogger,
		pgClient,
	)

	rows, err := client.CalculateRouteShapesFromTrips(context.TODO(), []int{55})
	assert.Error(t, err, fmt.Errorf(
		"no route shapes calculated, zero result from the query, route ids: %v",
		[]int{55},
	))
	assert.Nil(t, rows)
}

func TestClientRouteShapesCalculationSuccess(t *testing.T) {
	pgClient := sqlx.MustConnect("postgres", dbConnectionString)

	defer func() {
		err := pgClient.Close()
		if err != nil {
			panic(err)
		}
	}()

	client := postgres.NewClient(
		testLogger,
		pgClient,
	)

	rows, err := client.CalculateRouteShapesFromTrips(context.TODO(), nil)

	assert.NoError(t, err)

	// separate assertions for better readability

	assert.Len(t, rows, 4)
	assert.EqualValues(t, postgres.RouteShapeRow{
		RouteID:     "1",
		DirectionID: "0",
		Geometry:    "MULTILINESTRING((11.177994574032 43.7540034456826,11.1818857651156 43.7576500655217,11.1829380426644 43.7587548749205,11.1840608624419 43.75980842529,11.1848594421952 43.7604860551691,11.1855339940963 43.7610076758167,11.1860499965188 43.761331898126,11.1868382144634 43.7617132129471,11.1873273987129 43.7619086551052,11.1878210975953 43.7620758973506,11.1882463065166 43.76218823044,11.1895614551322 43.7624438548454,11.190557720936 43.762620124284,11.1940014338367 43.7631659259421,11.1946814807963 43.7633218052817,11.1955263920893 43.7635562502234,11.1973120342411 43.7640874815961,11.1982340433633 43.7643019341783,11.1988784364087 43.7644009839922,11.2029112370371 43.7647347558222,11.2042586615245 43.7648681866234,11.204876489998 43.7649462992771,11.2053689526242 43.7650358322987,11.2057279592949 43.7651279399934,11.2060712200719 43.7652563646967,11.2064979900819 43.7654597053688,11.2075812320431 43.7660735260421,11.207865966174 43.7662570973659,11.2089823908133 43.7671808846038,11.2091961401921 43.7673343128143,11.2093713638592 43.7675380050327,11.2095301197609 43.7678005383244,11.2097928988724 43.7681465926551,11.2110825399611 43.7701024121056,11.2114519481164 43.7707705250091,11.2121067310381 43.7718067566868,11.2129382952303 43.7731979235381,11.2207616989365 43.7741046364516,11.2217739779617 43.7742592487158,11.2281436286732 43.7750559807636,11.2287328761554 43.7752913852132,11.228823815445 43.7754703910433,11.2299580979638 43.7769604556257,11.2303639550765 43.7772500223541,11.2305022801278 43.7773060241118,11.2308420934115 43.7776188330743,11.2337362258162 43.7766709334981,11.2342864109139 43.7764504097053,11.2343178412147 43.7765231564892,11.2348164359821 43.7768568870065,11.235007111712 43.7768610078087,11.2353223059713 43.7768323745028,11.2354600577906 43.7767847465101,11.2356364140973 43.7766914303877,11.2357475375856 43.7765948974363,11.235903379757 43.7763671849423,11.2361162612546 43.7761507870032,11.2363782885149 43.7759799933843,11.2365513167473 43.7758883716945,11.2368134052013 43.7763190094278,11.2372657407841 43.7768783259783,11.2387704391287 43.7775039908911,11.2396271883452 43.7776961334714,11.239915173554 43.7777691189584,11.2400221116085 43.7778097079541,11.2401772026733 43.777894285049,11.2403473713329 43.778050458181,11.2405565620351 43.7781721663669,11.2408813172824 43.7783163201569,11.2414380223561 43.7785233544968,11.2425126264799 43.7788670293795,11.2427633365794 43.7789070028338,11.2428061983289 43.7788071218412,11.2429746546776 43.7786597585683,11.2457868995703 43.7773125061565),(11.2457868995703 43.7773125061565,11.2448850513217 43.7793200035566,11.2447223553787 43.7795267396988,11.2446267465947 43.7798472816385,11.2442309496473 43.7806023474908,11.2437692017822 43.7810480869763,11.2428026031642 43.7815352496218,11.2426341385906 43.7816826127543,11.243033896218 43.7820709176269,11.2429941593309 43.7823320961152,11.2429098855284 43.7824935139182,11.242968584867 43.7825673917373,11.242857697922 43.7829252065711,11.2427230210779 43.7832699831065,11.2419395835533 43.7851670892525,11.2418600934433 43.7853262102035,11.2418036284286 43.7853948444233,11.2414674060496 43.7857315213664,11.2406104298381 43.7864002870289,11.2403614961279 43.7864593183036,11.2402495050358 43.7865026678154,11.240181899393 43.7865610969741,11.2344656420353 43.7887438511516,11.2325059051144 43.789475434954,11.2313466934923 43.7890929269014,11.2310670899638 43.7890938862572,11.2301947489065 43.7891559374474,11.2298272997886 43.7891546675515,11.2295703857898 43.7891186058518,11.2265917440873 43.7906532725922,11.2256632085267 43.7910855271065,11.2243518777262 43.7916692839841,11.2221650537273 43.792693223077,11.22065432202 43.7933483421875,11.219819899861 43.7936391657051,11.2188594327708 43.7938829123028,11.2185349520241 43.7940152630715,11.2169510092403 43.794793301869,11.2170254232131 43.7948728889936,11.2137090227253 43.7964497876596,11.2128277116364 43.7969035312167,11.212837525771 43.7969168463211,11.2117610057178 43.7974283836681,11.2116817028908 43.7975109503203,11.2116598483957 43.7975924047022,11.2116733731394 43.7976641706463,11.2120223051461 43.7980670726113,11.2121493419245 43.79814114216,11.2122947287639 43.7982913857178,11.2120901069533 43.7983043517319,11.211759493246 43.7983557707201,11.2115379699863 43.7984261336099,11.2114399420741 43.7984848265074,11.2112939507918 43.7985403510461,11.2103334748002 43.798829043437,11.2078468782521 43.7997009470498,11.206999034858 43.8000954772232,11.2067101445174 43.8002676223688,11.2066117400455 43.8003607300854,11.2065034013668 43.8003666925197,11.2056771385359 43.8007654763569,11.2055670868719 43.8008049658251,11.205124810252 43.8009472054209,11.205009821172 43.800970763218,11.2047769381204 43.801013253099,11.2043931243745 43.8010449494741,11.2042866482538 43.8012140448273,11.2034410659152 43.8020335056091,11.2027273368536 43.8027753571884,11.2024080021208 43.802969165543,11.2021908961734 43.8030785344841,11.2017455024514 43.8033738166721,11.2016890982705 43.803472861519,11.2017324256314 43.8036193056553),(11.2457868995703 43.7773125061565,11.246200597123 43.7767461559909),(11.246200597123 43.7767461559909,11.2464988002758 43.7764373480361),(11.2464988002758 43.7764373480361,11.2472057017082 43.7758671720458),(11.2472057017082 43.7758671720458,11.2471717772092 43.775430552857),(11.2471717772092 43.775430552857,11.2496350362712 43.7753208437076),(11.2496350362712 43.7753208437076,11.2497673330381 43.7752565930201),(11.2497673330381 43.7752565930201,11.2498526996774 43.7752549117913),(11.2498526996774 43.7752549117913,11.2499878038398 43.7753014093176),(11.2499878038398 43.7753014093176,11.2500714736739 43.7750274578611),(11.2500714736739 43.7750274578611,11.2501984517261 43.7748371341095))",
		Centroid:    "POINT(11.218455974856017 43.778450894351245)",
	}, rows[0])

	assert.EqualValues(t, postgres.RouteShapeRow{
		RouteID:     "1",
		DirectionID: "1",
		Geometry:    "MULTILINESTRING((11.1818857651156 43.7576500655217,11.1774931577199 43.7535335016634,11.1769780227274 43.7537096530936,11.1766736582563 43.7539401018455,11.176239126542 43.7541599644981,11.175183489081 43.7547422594817),(11.1818857651156 43.7576500655217,11.177994574032 43.7540034456826),(11.1829380426644 43.7587548749205,11.1818857651156 43.7576500655217),(11.1840608624419 43.75980842529,11.1829380426644 43.7587548749205),(11.1848594421952 43.7604860551691,11.1840608624419 43.75980842529),(11.1855339940963 43.7610076758167,11.1848594421952 43.7604860551691),(11.1860499965188 43.761331898126,11.1855339940963 43.7610076758167),(11.1868382144634 43.7617132129471,11.1864323837762 43.761524204993,11.1860499965188 43.761331898126),(11.1868382144634 43.7617132129471,11.1860499965188 43.761331898126),(11.1873273987129 43.7619086551052,11.1868382144634 43.7617132129471),(11.1878210975953 43.7620758973506,11.1873273987129 43.7619086551052),(11.1882463065166 43.76218823044,11.1878210975953 43.7620758973506),(11.1895614551322 43.7624438548454,11.1882463065166 43.76218823044),(11.190557720936 43.762620124284,11.1895614551322 43.7624438548454),(11.1940014338367 43.7631659259421,11.190557720936 43.762620124284),(11.1946814807963 43.7633218052817,11.1940014338367 43.7631659259421),(11.1955263920893 43.7635562502234,11.1946814807963 43.7633218052817),(11.1973120342411 43.7640874815961,11.1955263920893 43.7635562502234),(11.1982340433633 43.7643019341783,11.1973120342411 43.7640874815961),(11.1988784364087 43.7644009839922,11.1982340433633 43.7643019341783),(11.2029112370371 43.7647347558222,11.1988784364087 43.7644009839922),(11.2017324256314 43.8036193056553,11.2016890982705 43.803472861519,11.2017455024514 43.8033738166721,11.2021908961734 43.8030785344841,11.2024080021208 43.802969165543,11.2027273368536 43.8027753571884,11.2034410659152 43.8020335056091,11.2042866482538 43.8012140448273,11.2043931243745 43.8010449494741,11.2047769381204 43.801013253099,11.205009821172 43.800970763218,11.205124810252 43.8009472054209,11.2055670868719 43.8008049658251,11.2056771385359 43.8007654763569,11.2065034013668 43.8003666925197,11.2066117400455 43.8003607300854,11.2067101445174 43.8002676223688,11.206999034858 43.8000954772232,11.2078468782521 43.7997009470498,11.2103334748002 43.798829043437,11.2112939507918 43.7985403510461,11.2114399420741 43.7984848265074,11.2115379699863 43.7984261336099,11.211759493246 43.7983557707201,11.2120901069533 43.7983043517319,11.2122947287639 43.7982913857178,11.2121493419245 43.79814114216,11.2120223051461 43.7980670726113,11.2116733731394 43.7976641706463,11.2116598483957 43.7975924047022,11.2116817028908 43.7975109503203,11.2117610057178 43.7974283836681,11.212837525771 43.7969168463211,11.2128277116364 43.7969035312167,11.2137090227253 43.7964497876596,11.2170254232131 43.7948728889936,11.2169510092403 43.794793301869,11.2185349520241 43.7940152630715,11.2188594327708 43.7938829123028,11.219819899861 43.7936391657051,11.22065432202 43.7933483421875,11.2221650537273 43.792693223077,11.2243518777262 43.7916692839841,11.2256632085267 43.7910855271065,11.2265917440873 43.7906532725922,11.2295703857898 43.7891186058518,11.2298272997886 43.7891546675515,11.2301947489065 43.7891559374474,11.2310670899638 43.7890938862572,11.2313466934923 43.7890929269014,11.2325059051144 43.789475434954,11.2344656420353 43.7887438511516,11.240181899393 43.7865610969741,11.2402495050358 43.7865026678154,11.2403614961279 43.7864593183036,11.2406104298381 43.7864002870289,11.2414674060496 43.7857315213664,11.2418036284286 43.7853948444233,11.2418600934433 43.7853262102035,11.2419395835533 43.7851670892525,11.2427230210779 43.7832699831065,11.242857697922 43.7829252065711,11.242968584867 43.7825673917373,11.2429098855284 43.7824935139182,11.2429941593309 43.7823320961152,11.243033896218 43.7820709176269,11.2426341385906 43.7816826127543,11.2428026031642 43.7815352496218,11.2437692017822 43.7810480869763,11.2442309496473 43.7806023474908,11.2446267465947 43.7798472816385,11.2447223553787 43.7795267396988,11.244885053171 43.7793200040509,11.2457868995703 43.7773125061565),(11.2042586615245 43.7648681866234,11.2029112370371 43.7647347558222),(11.204876489998 43.7649462992771,11.2042586615245 43.7648681866234),(11.205308011613319 43.76502475281256,11.2052484698669 43.7650111475389,11.204876489998 43.7649462992771),(11.205308011613319 43.76502475281256,11.204876489998 43.7649462992771),(11.205470350572796 43.76506184722281,11.2053689526242 43.7650358322987,11.205308011613319 43.76502475281256),(11.205470350572796 43.76506184722281,11.205308011613319 43.76502475281256),(11.20565294283284 43.76510869357279,11.2056118009784 43.7650941686039,11.205470350572796 43.76506184722281),(11.20565294283284 43.76510869357279,11.205470350572796 43.76506184722281),(11.2060712200719 43.7652563646967,11.2057279592949 43.7651279399934,11.20565294283284 43.76510869357279),(11.2060712200719 43.7652563646967,11.20565294283284 43.76510869357279),(11.206384152980842 43.76540546605465,11.2062907293519 43.7653556660784,11.2060712200719 43.7652563646967),(11.206384152980842 43.76540546605465,11.2060712200719 43.7652563646967),(11.206689744174641 43.76556836311151,11.2064979900819 43.7654597053688,11.206384152980842 43.76540546605465),(11.206689744174641 43.76556836311151,11.206384152980842 43.76540546605465),(11.207098445445938 43.76579995429659,11.2067960227713 43.7656250154986,11.206689744174641 43.76556836311151),(11.207098445445938 43.76579995429659,11.206689744174641 43.76556836311151),(11.207667265301495 43.766128992659624,11.2075812320431 43.7660735260421,11.207098445445938 43.76579995429659),(11.207667265301495 43.766128992659624,11.207098445445938 43.76579995429659),(11.207830947226268 43.76623452025309,11.2077750273522 43.7661913284737,11.207667265301495 43.766128992659624),(11.207830947226268 43.76623452025309,11.207667265301495 43.766128992659624),(11.20794716292331 43.76632428373758,11.207865966174 43.7662570973659,11.207830947226268 43.76623452025309),(11.20794716292331 43.76632428373758,11.207830947226268 43.76623452025309),(11.2089823908133 43.7671808846038,11.2081218789193 43.7664592320879,11.20794716292331 43.76632428373758),(11.2089823908133 43.7671808846038,11.20794716292331 43.76632428373758),(11.2091961401921 43.7673343128143,11.2089823908133 43.7671808846038),(11.2093713638592 43.7675380050327,11.2092853037604 43.7674361289493,11.2091961401921 43.7673343128143),(11.2093713638592 43.7675380050327,11.2091961401921 43.7673343128143),(11.2095301197609 43.7678005383244,11.2093713638592 43.7675380050327),(11.2097928988724 43.7681465926551,11.2095301197609 43.7678005383244),(11.2110825399611 43.7701024121056,11.2097928988724 43.7681465926551),(11.2114519481164 43.7707705250091,11.2110825399611 43.7701024121056),(11.2121067310381 43.7718067566868,11.2114519481164 43.7707705250091),(11.2129382952303 43.7731979235381,11.2121067310381 43.7718067566868),(11.2207616989365 43.7741046364516,11.2129382952303 43.7731979235381),(11.2217739779617 43.7742592487158,11.2207616989365 43.7741046364516),(11.2281436286732 43.7750559807636,11.2217739779617 43.7742592487158),(11.2287328761554 43.7752913852132,11.2281436286732 43.7750559807636),(11.228823815445 43.7754703910433,11.2287328761554 43.7752913852132),(11.2299580979638 43.7769604556257,11.228823815445 43.7754703910433),(11.2303639550765 43.7772500223541,11.2299580979638 43.7769604556257),(11.2305022801278 43.7773060241118,11.2303639550765 43.7772500223541),(11.2308420934115 43.7776188330743,11.2305022801278 43.7773060241118),(11.2337362258162 43.7766709334981,11.2308420934115 43.7776188330743),(11.2342864109139 43.7764504097053,11.2337362258162 43.7766709334981),(11.2343178412147 43.7765231564892,11.2342864109139 43.7764504097053),(11.2348164359821 43.7768568870065,11.2343178412147 43.7765231564892),(11.235007111712 43.7768610078087,11.2348164359821 43.7768568870065),(11.2353223059713 43.7768323745028,11.235007111712 43.7768610078087),(11.2354600577906 43.7767847465101,11.2353223059713 43.7768323745028),(11.2356364140973 43.7766914303877,11.2354600577906 43.7767847465101),(11.2357475375856 43.7765948974363,11.2356364140973 43.7766914303877),(11.235903379757 43.7763671849423,11.2357475375856 43.7765948974363),(11.2361162612546 43.7761507870032,11.235903379757 43.7763671849423),(11.2363782885149 43.7759799933843,11.2361162612546 43.7761507870032),(11.2365513167473 43.7758883716945,11.2363782885149 43.7759799933843),(11.2368134052013 43.7763190094278,11.2365513167473 43.7758883716945),(11.2372657407841 43.7768783259783,11.2368134052013 43.7763190094278),(11.2387704391287 43.7775039908911,11.2372657407841 43.7768783259783),(11.2396271883452 43.7776961334714,11.2387704391287 43.7775039908911),(11.239915173554 43.7777691189584,11.2396271883452 43.7776961334714),(11.2400221116085 43.7778097079541,11.239915173554 43.7777691189584),(11.2401772026733 43.777894285049,11.2400221116085 43.7778097079541),(11.2403473713329 43.778050458181,11.2401772026733 43.777894285049),(11.2405565620351 43.7781721663669,11.2403473713329 43.778050458181),(11.2408813172824 43.7783163201569,11.2405565620351 43.7781721663669),(11.2414380223561 43.7785233544968,11.2408813172824 43.7783163201569),(11.2425126264799 43.7788670293795,11.2414380223561 43.7785233544968),(11.2427633365794 43.7789070028338,11.2425126264799 43.7788670293795),(11.2428061983289 43.7788071218412,11.2427633365794 43.7789070028338),(11.2429746546776 43.7786597585683,11.2428061983289 43.7788071218412),(11.2457868995703 43.7773125061565,11.2429746546776 43.7786597585683),(11.246200597123 43.7767461559909,11.2457868995703 43.7773125061565),(11.2464988002758 43.7764373480361,11.246200597123 43.7767461559909),(11.2472057017082 43.7758671720458,11.2464988002758 43.7764373480361),(11.2471717772092 43.775430552857,11.2472057017082 43.7758671720458),(11.2496350362712 43.7753208437076,11.2471717772092 43.775430552857),(11.2497673330381 43.7752565930201,11.2496350362712 43.7753208437076),(11.2498526996774 43.7752549117913,11.2497673330381 43.7752565930201),(11.2499878038398 43.7753014093176,11.2498526996774 43.7752549117913),(11.2500714736739 43.7750274578611,11.2499878038398 43.7753014093176),(11.2501984517261 43.7748371341095,11.2500714736739 43.7750274578611))",
		Centroid:    "POINT(11.215689952505773 43.776650567538034)",
	}, rows[1])

	assert.EqualValues(t, postgres.RouteShapeRow{
		RouteID:     "2",
		DirectionID: "0",
		Geometry:    "MULTILINESTRING((11.1751791991596 43.7547446753329,11.176239126542 43.7541599644981,11.1766736582563 43.7539401018455,11.1769780227274 43.7537096530936,11.1774931577199 43.7535335016634,11.1818857651156 43.7576500655217),(11.177994574032 43.7540034456826,11.1818857651156 43.7576500655217),(11.1818857651156 43.7576500655217,11.1829380426644 43.7587548749205),(11.1829380426644 43.7587548749205,11.1840608624419 43.75980842529),(11.1840608624419 43.75980842529,11.1848594421952 43.7604860551691),(11.1848594421952 43.7604860551691,11.1855339940963 43.7610076758167),(11.1855339940963 43.7610076758167,11.1860499965188 43.761331898126),(11.1860499965188 43.761331898126,11.1864323837762 43.761524204993,11.1868382144634 43.7617132129471),(11.1860499965188 43.761331898126,11.1868382144634 43.7617132129471),(11.1868382144634 43.7617132129471,11.1873273987129 43.7619086551052),(11.1873273987129 43.7619086551052,11.1878210975953 43.7620758973506),(11.1878210975953 43.7620758973506,11.1882463065166 43.76218823044),(11.1882463065166 43.76218823044,11.1895614551322 43.7624438548454),(11.1895614551322 43.7624438548454,11.190557720936 43.762620124284),(11.190557720936 43.762620124284,11.1940014338367 43.7631659259421),(11.1940014338367 43.7631659259421,11.1946814807963 43.7633218052817),(11.1946814807963 43.7633218052817,11.1955263920893 43.7635562502234),(11.1955263920893 43.7635562502234,11.1973120342411 43.7640874815961),(11.1973120342411 43.7640874815961,11.1982340433633 43.7643019341783),(11.1982340433633 43.7643019341783,11.1988784364087 43.7644009839922),(11.1988784364087 43.7644009839922,11.2029112370371 43.7647347558222),(11.2029112370371 43.7647347558222,11.2042586615245 43.7648681866234),(11.2042586615245 43.7648681866234,11.204876489998 43.7649462992771),(11.204876489998 43.7649462992771,11.2052484698669 43.7650111475389,11.205308011613319 43.76502475281256),(11.204876489998 43.7649462992771,11.205308011613319 43.76502475281256),(11.205308011613319 43.76502475281256,11.2053689526242 43.7650358322987,11.205470350572796 43.76506184722281),(11.205308011613319 43.76502475281256,11.205470350572796 43.76506184722281),(11.205470350572796 43.76506184722281,11.2056118009784 43.7650941686039,11.20565294283284 43.76510869357279),(11.205470350572796 43.76506184722281,11.20565294283284 43.76510869357279),(11.20565294283284 43.76510869357279,11.2057279592949 43.7651279399934,11.2060712200719 43.7652563646967),(11.20565294283284 43.76510869357279,11.2060712200719 43.7652563646967),(11.2060712200719 43.7652563646967,11.2062907293519 43.7653556660784,11.206384152980842 43.76540546605465),(11.2060712200719 43.7652563646967,11.206384152980842 43.76540546605465),(11.206384152980842 43.76540546605465,11.2064979900819 43.7654597053688,11.206689744174641 43.76556836311151),(11.206384152980842 43.76540546605465,11.206689744174641 43.76556836311151),(11.206689744174641 43.76556836311151,11.2067960227713 43.7656250154986,11.207098445445938 43.76579995429659),(11.206689744174641 43.76556836311151,11.207098445445938 43.76579995429659),(11.207098445445938 43.76579995429659,11.2075812320431 43.7660735260421,11.207667265301495 43.766128992659624),(11.207098445445938 43.76579995429659,11.207667265301495 43.766128992659624),(11.207667265301495 43.766128992659624,11.2077750273522 43.7661913284737,11.207830947226268 43.76623452025309),(11.207667265301495 43.766128992659624,11.207830947226268 43.76623452025309),(11.207830947226268 43.76623452025309,11.207865966174 43.7662570973659,11.20794716292331 43.76632428373758),(11.207830947226268 43.76623452025309,11.20794716292331 43.76632428373758),(11.20794716292331 43.76632428373758,11.2081218789193 43.7664592320879,11.2089823908133 43.7671808846038),(11.20794716292331 43.76632428373758,11.2089823908133 43.7671808846038),(11.2089823908133 43.7671808846038,11.2091961401921 43.7673343128143),(11.2091961401921 43.7673343128143,11.2092853037604 43.7674361289493,11.2093713638592 43.7675380050327),(11.2091961401921 43.7673343128143,11.2093713638592 43.7675380050327),(11.2093713638592 43.7675380050327,11.2095301197609 43.7678005383244),(11.2095301197609 43.7678005383244,11.2097928988724 43.7681465926551),(11.2097928988724 43.7681465926551,11.2110825399611 43.7701024121056),(11.2110825399611 43.7701024121056,11.2114519481164 43.7707705250091),(11.2114519481164 43.7707705250091,11.2121067310381 43.7718067566868),(11.2121067310381 43.7718067566868,11.2129382952303 43.7731979235381),(11.2129382952303 43.7731979235381,11.2207616989365 43.7741046364516),(11.2207616989365 43.7741046364516,11.2217739779617 43.7742592487158),(11.2217739779617 43.7742592487158,11.2281436286732 43.7750559807636),(11.2281436286732 43.7750559807636,11.2287328761554 43.7752913852132),(11.2287328761554 43.7752913852132,11.228823815445 43.7754703910433),(11.228823815445 43.7754703910433,11.2299580979638 43.7769604556257),(11.2299580979638 43.7769604556257,11.2303639550765 43.7772500223541),(11.2303639550765 43.7772500223541,11.2305022801278 43.7773060241118),(11.2305022801278 43.7773060241118,11.2308420934115 43.7776188330743),(11.2308420934115 43.7776188330743,11.2337362258162 43.7766709334981),(11.2337362258162 43.7766709334981,11.2342864109139 43.7764504097053),(11.2342864109139 43.7764504097053,11.2343178412147 43.7765231564892),(11.2343178412147 43.7765231564892,11.2348164359821 43.7768568870065),(11.2348164359821 43.7768568870065,11.235007111712 43.7768610078087),(11.235007111712 43.7768610078087,11.2353223059713 43.7768323745028),(11.2353223059713 43.7768323745028,11.2354600577906 43.7767847465101),(11.2354600577906 43.7767847465101,11.2356364140973 43.7766914303877),(11.2356364140973 43.7766914303877,11.2357475375856 43.7765948974363),(11.2357475375856 43.7765948974363,11.235903379757 43.7763671849423),(11.235903379757 43.7763671849423,11.2361162612546 43.7761507870032),(11.2361162612546 43.7761507870032,11.2363782885149 43.7759799933843),(11.2363782885149 43.7759799933843,11.2365513167473 43.7758883716945),(11.2365513167473 43.7758883716945,11.2368134052013 43.7763190094278),(11.2368134052013 43.7763190094278,11.2372657407841 43.7768783259783),(11.2372657407841 43.7768783259783,11.2387704391287 43.7775039908911),(11.2387704391287 43.7775039908911,11.2396271883452 43.7776961334714),(11.2396271883452 43.7776961334714,11.239915173554 43.7777691189584),(11.239915173554 43.7777691189584,11.2400221116085 43.7778097079541),(11.2400221116085 43.7778097079541,11.2401772026733 43.777894285049),(11.2401772026733 43.777894285049,11.2403473713329 43.778050458181),(11.2403473713329 43.778050458181,11.2405565620351 43.7781721663669),(11.2405565620351 43.7781721663669,11.2408813172824 43.7783163201569),(11.2407075740416 43.7964275986966,11.2408425726331 43.7965059818268),(11.2439752805611 43.7942746511116,11.2407075740416 43.7964275986966),(11.2408425726331 43.7965059818268,11.2410199577633 43.7966375540798),(11.2408813172824 43.7783163201569,11.2414380223561 43.7785233544968),(11.2410199577633 43.7966375540798,11.2419319973327 43.7973849529349),(11.2414380223561 43.7785233544968,11.2425126264799 43.7788670293795),(11.2419319973327 43.7973849529349,11.2426224812142 43.7980241508245),(11.2425126264799 43.7788670293795,11.2427633365794 43.7789070028338),(11.2426224812142 43.7980241508245,11.2431405784455 43.7985406811322),(11.2427633365794 43.7789070028338,11.2428061983289 43.7788071218412),(11.2428061983289 43.7788071218412,11.2429746546776 43.7786597585683),(11.2429746546776 43.7786597585683,11.2457868995703 43.7773125061565),(11.2431405784455 43.7985406811322,11.2438848378973 43.7994534208803),(11.2438848378973 43.7994534208803,11.2442263253745 43.7999419029903),(11.2442184662053 43.7941303200287,11.2439752805611 43.7942746511116),(11.2443739636566 43.7940507353005,11.2442184662053 43.7941303200287),(11.2442263253745 43.7999419029903,11.2443786073694 43.8001910081006),(11.2447909565905 43.7939840195566,11.2443739636566 43.7940507353005),(11.2443786073694 43.8001910081006,11.2449774814316 43.8012866659737),(11.2451377494736 43.7939096782697,11.2447909565905 43.7939840195566),(11.2449774814316 43.8012866659737,11.2453967387149 43.8022327924249),(11.2452707977509 43.7938530424714,11.2451377494736 43.7939096782697),(11.2453967930857 43.7937740364757,11.2452707977509 43.7938530424714),(11.2453967387149 43.8022327924249,11.2457781616932 43.8031933799288),(11.2455696039988 43.7936175805732,11.2453967930857 43.7937740364757),(11.245783392132 43.7933522771795,11.2455696039988 43.7936175805732),(11.2457781616932 43.8031933799288,11.2458711120442 43.8032068781442),(11.2462307103852 43.7927267445295,11.245783392132 43.7933522771795),(11.2457868995703 43.7773125061565,11.246200597123 43.7767461559909),(11.2458711120442 43.8032068781442,11.2459532315838 43.8032400426234),(11.2459532315838 43.8032400426234,11.2459949098527 43.8033155269569),(11.2459949098527 43.8033155269569,11.2461293955949 43.8032898002694),(11.2461293955949 43.8032898002694,11.2466173388274 43.8032486928401),(11.246200597123 43.7767461559909,11.2464988002758 43.7764373480361),(11.2463256956116 43.7926855345789,11.2462307103852 43.7927267445295),(11.2476193334595 43.7908838444869,11.2463256956116 43.7926855345789),(11.2464988002758 43.7764373480361,11.2472057017082 43.7758671720458),(11.2466173388274 43.8032486928401,11.2469652378864 43.8031857248547),(11.2472057017082 43.7758671720458,11.2471717772092 43.775430552857),(11.2471717772092 43.775430552857,11.2491203164583 43.7753436760426),(11.2483287951336 43.7898448742868,11.2476193334595 43.7908838444869),(11.2483148842904 43.7898136359126,11.2483287951336 43.7898448742868),(11.2498926854046 43.7875947268939,11.2483148842904 43.7898136359126),(11.2483469836587 43.7799316834114,11.2483638545727 43.7799673653166),(11.2483672646587 43.7798097371607,11.2483469836587 43.7799316834114),(11.2483638545727 43.7799673653166,11.2484098794077 43.7799934699253),(11.2485828187393 43.7792202677398,11.2483672646587 43.7798097371607),(11.2484098794077 43.7799934699253,11.2488560934795 43.7801287424373),(11.2488848852267 43.7782869621891,11.2485828187393 43.7792202677398),(11.2488560934795 43.7801287424373,11.2488641238673 43.7802186192256),(11.2488641238673 43.7802186192256,11.249512564656 43.7804444448202),(11.2490830322859 43.7774207969696,11.2488848852267 43.7782869621891),(11.2496746302222 43.7759034941532,11.2490830322859 43.7774207969696),(11.2491203164583 43.7753436760426,11.2491734540483 43.7756082101612),(11.2491734540483 43.7756082101612,11.24932489043 43.7757583166821),(11.24932489043 43.7757583166821,11.2496746302222 43.7759034941532),(11.249512564656 43.7804444448202,11.250579018013 43.7807160525451),(11.2503047827145 43.7866952654856,11.2498926854046 43.7875947268939),(11.2512665504503 43.7847135582929,11.2503047827145 43.7866952654856),(11.250579018013 43.7807160525451,11.2509450998455 43.7807853693754),(11.2509450998455 43.7807853693754,11.2517010874857 43.7808965194585),(11.2513550677797 43.7846539880838,11.2512665504503 43.7847135582929),(11.2515780110373 43.7842045069568,11.2513550677797 43.7846539880838),(11.2518237055183 43.7838102041414,11.2515780110373 43.7842045069568),(11.2517010874857 43.7808965194585,11.252092112113 43.7810508740662),(11.2519304696673 43.7836731819536,11.2518237055183 43.7838102041414),(11.2520210952051 43.7833405372284,11.2519304696673 43.7836731819536),(11.2521137476003 43.7831572627913,11.2520210952051 43.7833405372284),(11.252092112113 43.7810508740662,11.2522282089921 43.7812417658542),(11.2522171670475 43.7830156698886,11.2521137476003 43.7831572627913),(11.2524141095824 43.782593124936,11.2522171670475 43.7830156698886),(11.2522282089921 43.7812417658542,11.2525876355768 43.7816713480001),(11.2524680879277 43.7823759770181,11.2524141095824 43.782593124936),(11.2526452485903 43.7818817938114,11.2524680879277 43.7823759770181),(11.2525876355768 43.7816713480001,11.2526452485903 43.7818817938114))",
		Centroid:    "POINT(11.220622696807853 43.77426355093984)",
	}, rows[2])

	assert.EqualValues(t, postgres.RouteShapeRow{
		RouteID:     "2",
		DirectionID: "1",
		Geometry:    "LINESTRING(11.2469652378864 43.8031857248547,11.2466173388274 43.8032486928401,11.2461293955949 43.8032898002694,11.2459949098527 43.8033155269569,11.2459532315838 43.8032400426234,11.2458711120442 43.8032068781442,11.2457781616932 43.8031933799288,11.2453967387149 43.8022327924249,11.2449774814316 43.8012866659737,11.2443786073694 43.8001910081006,11.2442263253745 43.7999419029903,11.2438848378973 43.7994534208803,11.2431405784455 43.7985406811322,11.2426224812142 43.7980241508245,11.2420107685036 43.7974554346919,11.2408425726331 43.7965059818268,11.2406282621937 43.7963841370329,11.2403134092735 43.7962327515332,11.2407930490674 43.7952644746215,11.2409976265975 43.7949183296561,11.2412892086354 43.794322881875,11.2414799616299 43.7938149447658,11.2446502494597 43.7929648905379,11.2460096298563 43.7927473319999,11.2462307103852 43.7927267445295,11.2461848201074 43.7926582874196,11.2475086606563 43.7908504889048,11.2482436145971 43.7898105368478,11.2483163586467 43.7898125906125,11.2498926854046 43.7875947268939,11.2503047827145 43.7866952654856,11.2512665504503 43.7847135582929,11.2513550677797 43.7846539880838,11.2515780110373 43.7842045069568,11.2518237055183 43.7838102041414,11.2519304696673 43.7836731819536,11.2520210952051 43.7833405372284,11.2521137476003 43.7831572627913,11.2522171670475 43.7830156698886,11.2524141095824 43.782593124936,11.2524680879277 43.7823759770181,11.2526452485903 43.7818817938114,11.2525876355768 43.7816713480001,11.2522282089921 43.7812417658542,11.252092112113 43.7810508740662,11.2517010874857 43.7808965194585,11.2509450998455 43.7807853693754,11.250579018013 43.7807160525451,11.249512564656 43.7804444448202,11.2488641257079 43.7802186198666,11.2488560934795 43.7801287424373,11.2484098794077 43.7799934699253,11.2483638545727 43.7799673653166,11.2483469836587 43.7799316834114,11.2483672646587 43.7798097371607,11.2485828187393 43.7792202677398,11.2488848852267 43.7782869621891,11.2490830322859 43.7774207969696,11.2496746302222 43.7759034941532,11.24932489043 43.7757583166821,11.2491734540483 43.7756082101612,11.2491203164583 43.7753436760426,11.2471717772092 43.775430552857,11.2472057017082 43.7758671720458,11.2464988002758 43.7764373480361,11.246200597123 43.7767461559909,11.2457868995703 43.7773125061565,11.2429746546776 43.7786597585683,11.2428061983289 43.7788071218412,11.2427633365794 43.7789070028338,11.2425126264799 43.7788670293795,11.2414380223561 43.7785233544968,11.2408813172824 43.7783163201569,11.2405565620351 43.7781721663669,11.2403473713329 43.778050458181,11.2401772026733 43.777894285049,11.2400221116085 43.7778097079541,11.239915173554 43.7777691189584,11.2396271883452 43.7776961334714,11.2387704391287 43.7775039908911,11.2372657407841 43.7768783259783,11.2368134052013 43.7763190094278,11.2365513167473 43.7758883716945,11.2363782885149 43.7759799933843,11.2361162612546 43.7761507870032,11.235903379757 43.7763671849423,11.2357475375856 43.7765948974363,11.2356364140973 43.7766914303877,11.2354600577906 43.7767847465101,11.2353223059713 43.7768323745028,11.235007111712 43.7768610078087,11.2348164359821 43.7768568870065,11.2343178412147 43.7765231564892,11.2342864109139 43.7764504097053,11.2337362258162 43.7766709334981,11.2308420934115 43.7776188330743,11.2305022801278 43.7773060241118,11.2303639550765 43.7772500223541,11.2299580979638 43.7769604556257,11.228823815445 43.7754703910433,11.2287328761554 43.7752913852132,11.2281436286732 43.7750559807636,11.2217739779617 43.7742592487158,11.2207616989365 43.7741046364516,11.2129382952303 43.7731979235381,11.2121067310381 43.7718067566868,11.2114519481164 43.7707705250091,11.2110825399611 43.7701024121056,11.2097928988724 43.7681465926551,11.2095301197609 43.7678005383244,11.2093713638592 43.7675380050327,11.2092853037604 43.7674361289493,11.2091961401921 43.7673343128143,11.2089823908133 43.7671808846038,11.2081218789193 43.7664592320879,11.2077750273522 43.7661913284737,11.2067960227713 43.7656250154986,11.2062907293519 43.7653556660784,11.2060712200719 43.7652563646967,11.2056118009784 43.7650941686039,11.2052484698669 43.7650111475389,11.204876489998 43.7649462992771,11.2042586615245 43.7648681866234,11.2029112370371 43.7647347558222,11.1988784364087 43.7644009839922,11.1982340433633 43.7643019341783,11.1973120342411 43.7640874815961,11.1955263920893 43.7635562502234,11.1946814807963 43.7633218052817,11.1940014338367 43.7631659259421,11.190557720936 43.762620124284,11.1895614551322 43.7624438548454,11.1882463065166 43.76218823044,11.1878210975953 43.7620758973506,11.1873273987129 43.7619086551052,11.1868382144634 43.7617132129471,11.1864323837762 43.761524204993,11.1860499965188 43.761331898126,11.1855339940963 43.7610076758167,11.1848594421952 43.7604860551691,11.1840608624419 43.75980842529,11.1829380426644 43.7587548749205,11.1818857651156 43.7576500655217,11.1774931577199 43.7535335016634,11.1769780227274 43.7537096530936,11.1766736582563 43.7539401018455,11.176239126542 43.7541599644981,11.175183489081 43.7547422594817)",
		Centroid:    "POINT(11.223279594575853 43.775618757072706)",
	}, rows[3])
}
