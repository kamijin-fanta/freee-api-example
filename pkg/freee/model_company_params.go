/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <pre><code>【重要】会計freee APIの新バージョンについて 2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br> 新しいAPIを利用するにはリクエストヘッダーに以下を指定します。 X-Api-Version: 2020-06-15<br> 指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br> 【重要】APIのバージョン指定をせずに利用し続ける場合 2020年12月に新しいバージョンのAPIに自動的に切り替わります。 詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br> 旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。 </code></pre>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package freee

import (
	"encoding/json"
)

// CompanyParams struct for CompanyParams
type CompanyParams struct {
	// 事業所の正式名称 (100文字以内)
	Name *string `json:"name,omitempty"`
	// 正式名称フリガナ (100文字以内)
	NameKana *string `json:"name_kana,omitempty"`
	// 担当者名 (50文字以内)
	ContactName *string `json:"contact_name,omitempty"`
	// 従業員数（0: 経営者のみ、1: 2~5人、2: 6~10人、3: 11~20人、13: 21~50人、14: 51~100人、15: 101~300人、18: 301~500人、16: 501~1,000人、17: 1,001人以上
	HeadCount *float32 `json:"head_count,omitempty"`
	// 法人番号 (半角数字13桁、法人のみ)
	CorporateNumber *string `json:"corporate_number,omitempty"`
	// 仕訳番号形式（not_used: 使用しない、digits: 数字（例：5091824）、alnum: 英数字（例：59J0P））
	TxnNumberFormat *string `json:"txn_number_format,omitempty"`
	// プライベート資金/役員資金（false: 使用しない、true: 使用する）
	PrivateSettlement *bool `json:"private_settlement,omitempty"`
	// 電話番号１
	Phone1 *string `json:"phone1,omitempty"`
	// 電話番号２
	Phone2 *string `json:"phone2,omitempty"`
	// FAX
	Fax *string `json:"fax,omitempty"`
	// 郵便番号
	Zipcode *string `json:"zipcode,omitempty"`
	// 都道府県コード（0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PrefectureCode *int32 `json:"prefecture_code,omitempty"`
	// 市区町村・番地
	StreetName1 *string `json:"street_name1,omitempty"`
	// 建物名・部屋番号など
	StreetName2 *string `json:"street_name2,omitempty"`
	// 請求書レイアウト * `default_classic` - レイアウト１/クラシック (デフォルト)  * `standard_classic` - レイアウト２/クラシック  * `envelope_classic` - 封筒１/クラシック  * `carried_forward_standard_classic` - レイアウト３（繰越金額欄あり）/クラシック  * `carried_forward_envelope_classic` - 封筒２（繰越金額欄あり）/クラシック  * `default_modern` - レイアウト１/モダン  * `standard_modern` - レイアウト２/モダン  * `envelope_modern` - 封筒/モダン
	InvoiceLayout *string `json:"invoice_layout,omitempty"`
	// 金額端数処理方法（0: 切り捨て、1: 切り上げ, 2: 四捨五入）
	AmountFraction *int32 `json:"amount_fraction,omitempty"`
	// 種別（agriculture_forestry_fisheries_ore: 農林水産業/鉱業、construction: 建設、manufacturing_processing: 製造/加工、it: IT、transportation_logistics: 運輸/物流、retail_wholesale: 小売/卸売、finance_insurance: 金融/保険、real_estate_rental: 不動産/レンタル、profession: 士業/学術/専門技術サービス、design_production: デザイン/制作、food: 飲食、leisure_entertainment: レジャー/娯楽、lifestyle: 生活関連サービス、education: 教育/学習支援、medical_welfare: 医療/福祉、other_services: その他サービス、other: その他）
	IndustryClass *string `json:"industry_class,omitempty"`
	// 業種（agriculture: 農業, forestry: 林業, fishing_industry: 漁業、水産養殖業, mining: 鉱業、採石業、砂利採取業, civil_contractors: 土木工事業, pavement: 舗装工事業, carpenter: とび、大工、左官等の建設工事業, renovation: リフォーム工事業, electrical_plumbing: 電気、管工事等の設備工事業, grocery: 食料品の製造加工業, machinery_manufacturing: 機械器具の製造加工業, printing: 印刷業, other_manufacturing: その他の製造加工業, software_development: 受託：ソフトウェア、アプリ開発業, system_development: 受託：システム開発業, survey_analysis: 受託：調査、分析等の情報処理業, server_management: 受託：サーバー運営管理, website_production: 受託：ウェブサイト制作, online_service_management: オンラインサービス運営業, online_advertising_agency: オンライン広告代理店業, online_advertising_planning_production: オンライン広告企画・制作業, online_media_management: オンラインメディア運営業, portal_site_management: ポータルサイト運営業, other_it_services: その他、IT サービス業, transport_delivery: 輸送業、配送業, delivery: バイク便等の配達業, other_transportation_logistics: その他の運輸業、物流業, other_wholesale: 卸売業：その他, clothing_wholesale_fiber: 卸売業：衣類卸売／繊維, food_wholesale: 卸売業：飲食料品, entrusted_development_wholesale: 卸売業：機械器具, online_shop: 小売業：無店舗　オンラインショップ, fashion_grocery_store: 小売業：店舗あり　ファッション、雑貨, food_store: 小売業：店舗あり　生鮮食品、飲食料品, entrusted_store: 小売業：店舗あり　機械、器具, other_store: 小売業：店舗あり　その他, financial_instruments_exchange: 金融業：金融商品取引, commodity_futures_investment_advisor: 金融業：商品先物取引、商品投資顧問, other_financial: 金融業：その他, brokerage_insurance: 保険業：仲介、代理, other_insurance: 保険業：その他, real_estate_developer: 不動産業：ディベロッパー, real_estate_brokerage: 不動産業：売買、仲介, rent_coin_parking_management: 不動産業：賃貸、コインパーキング、管理, rental_office_co_working_space: 不動産業：レンタルオフィス、コワーキングスペース, rental_lease: レンタル業、リース業, cpa_tax_accountant: 士業：公認会計士事務所、税理士事務所, law_office: 士業：法律事務所, judicial_and_administrative_scrivener: 士業：司法書士事務所／行政書士事務所, labor_consultant: 士業：社会保険労務士事務所, other_profession: 士業：その他, business_consultant: 経営コンサルタント, academic_research_development: 学術・開発研究機関, advertising_agency: 広告代理店, advertising_planning_production: 広告企画／制作, design_development: ソフトウェア、アプリ開発業（受託）, apparel_industry_design: 服飾デザイン業、工業デザイン業, website_design: ウェブサイト制作（受託）, advertising_planning_design: 広告企画／制作業, other_design: その他、デザイン／制作, restaurants_coffee_shops: レストラン、喫茶店等の飲食店業, sale_of_lunch: 弁当の販売業, bread_confectionery_manufacture_sale: パン、菓子等の製造販売業, delivery_catering_mobile_catering: デリバリー業、ケータリング業、移動販売業, hotel_inn: 宿泊業：ホテル、旅館, homestay: 宿泊業：民泊, travel_agency: 旅行代理店業, leisure_sports_facility_management: レジャー、スポーツ等の施設運営業, show_event_management: ショー、イベント等の興行、イベント運営業, barber: ビューティ、ヘルスケア業：床屋、理容室, beauty_salon: ビューティ、ヘルスケア業：美容室, spa_sand_bath_sauna: ビューティ、ヘルスケア業：スパ、砂風呂、サウナ等, este_ail_salon: ビューティ、ヘルスケア業：その他、エステサロン、ネイルサロン等, bridal_planning_introduce_wedding: 冠婚葬祭業：ブライダルプランニング、結婚式場紹介等, memorial_ceremony_funeral: 冠婚葬祭業：メモリアルセレモニー、葬儀等, moving: 引っ越し業, courier_industry: 宅配業, house_maid_cleaning_agency: 家事代行サービス業：無店舗　ハウスメイド、掃除代行等, re_tailoring_clothes: 家事代行サービス業：店舗あり　衣類修理、衣類仕立て直し等, training_institute_management: 研修所等の施設運営業, tutoring_school: 学習塾、進学塾等の教育・学習支援業, music_calligraphy_abacus_classroom: 音楽教室、書道教室、そろばん教室等のの教育・学習支援業, english_school: 英会話スクール等の語学学習支援業, tennis_yoga_judo_school: テニススクール、ヨガ教室、柔道場等のスポーツ指導、支援業, culture_school: その他、カルチャースクール等の教育・学習支援業, seminar_planning_management: セミナー等の企画、運営業, hospital_clinic: 医療業：病院、一般診療所、クリニック等, dental_clinic: 医療業：歯科診療所, other_medical_services: 医療業：その他、医療サービス等, nursery: 福祉業：保育所等、児童向け施設型サービス, nursing_home: 福祉業：老人ホーム等、老人向け施設型サービス, rehabilitation_support_services: 福祉業：療育支援サービス等、障害者等向け施設型サービス, other_welfare: 福祉業：その他、施設型福祉サービス, visit_welfare_service: 福祉業：訪問型福祉サービス, recruitment_temporary_staffing: 人材紹介業、人材派遣業, life_related_recruitment_temporary_staffing: 生活関連サービスの人材紹介業、人材派遣業, car_maintenance_car_repair: 自動車整備業、自動車修理業, machinery_equipment_maintenance_repair: 機械機器類の整備業、修理業, cleaning_maintenance_building_management: 清掃業、メンテナンス業、建物管理業, security: 警備業, other_services: その他のサービス業, npo: NPO, general_incorporated_association: 一般社団法人, general_incorporated_foundation: 一般財団法人, other_association: その他組織)
	IndustryCode *string `json:"industry_code,omitempty"`
	// 決済口座のデフォルト
	DefaultWalletAccountId *int32 `json:"default_wallet_account_id,omitempty"`
	FiscalYears *CompanyParamsFiscalYears `json:"fiscal_years,omitempty"`
}

// NewCompanyParams instantiates a new CompanyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyParams() *CompanyParams {
	this := CompanyParams{}
	return &this
}

// NewCompanyParamsWithDefaults instantiates a new CompanyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyParamsWithDefaults() *CompanyParams {
	this := CompanyParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyParams) SetName(v string) {
	o.Name = &v
}

// GetNameKana returns the NameKana field value if set, zero value otherwise.
func (o *CompanyParams) GetNameKana() string {
	if o == nil || o.NameKana == nil {
		var ret string
		return ret
	}
	return *o.NameKana
}

// GetNameKanaOk returns a tuple with the NameKana field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetNameKanaOk() (*string, bool) {
	if o == nil || o.NameKana == nil {
		return nil, false
	}
	return o.NameKana, true
}

// HasNameKana returns a boolean if a field has been set.
func (o *CompanyParams) HasNameKana() bool {
	if o != nil && o.NameKana != nil {
		return true
	}

	return false
}

// SetNameKana gets a reference to the given string and assigns it to the NameKana field.
func (o *CompanyParams) SetNameKana(v string) {
	o.NameKana = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *CompanyParams) GetContactName() string {
	if o == nil || o.ContactName == nil {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetContactNameOk() (*string, bool) {
	if o == nil || o.ContactName == nil {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *CompanyParams) HasContactName() bool {
	if o != nil && o.ContactName != nil {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *CompanyParams) SetContactName(v string) {
	o.ContactName = &v
}

// GetHeadCount returns the HeadCount field value if set, zero value otherwise.
func (o *CompanyParams) GetHeadCount() float32 {
	if o == nil || o.HeadCount == nil {
		var ret float32
		return ret
	}
	return *o.HeadCount
}

// GetHeadCountOk returns a tuple with the HeadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetHeadCountOk() (*float32, bool) {
	if o == nil || o.HeadCount == nil {
		return nil, false
	}
	return o.HeadCount, true
}

// HasHeadCount returns a boolean if a field has been set.
func (o *CompanyParams) HasHeadCount() bool {
	if o != nil && o.HeadCount != nil {
		return true
	}

	return false
}

// SetHeadCount gets a reference to the given float32 and assigns it to the HeadCount field.
func (o *CompanyParams) SetHeadCount(v float32) {
	o.HeadCount = &v
}

// GetCorporateNumber returns the CorporateNumber field value if set, zero value otherwise.
func (o *CompanyParams) GetCorporateNumber() string {
	if o == nil || o.CorporateNumber == nil {
		var ret string
		return ret
	}
	return *o.CorporateNumber
}

// GetCorporateNumberOk returns a tuple with the CorporateNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetCorporateNumberOk() (*string, bool) {
	if o == nil || o.CorporateNumber == nil {
		return nil, false
	}
	return o.CorporateNumber, true
}

// HasCorporateNumber returns a boolean if a field has been set.
func (o *CompanyParams) HasCorporateNumber() bool {
	if o != nil && o.CorporateNumber != nil {
		return true
	}

	return false
}

// SetCorporateNumber gets a reference to the given string and assigns it to the CorporateNumber field.
func (o *CompanyParams) SetCorporateNumber(v string) {
	o.CorporateNumber = &v
}

// GetTxnNumberFormat returns the TxnNumberFormat field value if set, zero value otherwise.
func (o *CompanyParams) GetTxnNumberFormat() string {
	if o == nil || o.TxnNumberFormat == nil {
		var ret string
		return ret
	}
	return *o.TxnNumberFormat
}

// GetTxnNumberFormatOk returns a tuple with the TxnNumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetTxnNumberFormatOk() (*string, bool) {
	if o == nil || o.TxnNumberFormat == nil {
		return nil, false
	}
	return o.TxnNumberFormat, true
}

// HasTxnNumberFormat returns a boolean if a field has been set.
func (o *CompanyParams) HasTxnNumberFormat() bool {
	if o != nil && o.TxnNumberFormat != nil {
		return true
	}

	return false
}

// SetTxnNumberFormat gets a reference to the given string and assigns it to the TxnNumberFormat field.
func (o *CompanyParams) SetTxnNumberFormat(v string) {
	o.TxnNumberFormat = &v
}

// GetPrivateSettlement returns the PrivateSettlement field value if set, zero value otherwise.
func (o *CompanyParams) GetPrivateSettlement() bool {
	if o == nil || o.PrivateSettlement == nil {
		var ret bool
		return ret
	}
	return *o.PrivateSettlement
}

// GetPrivateSettlementOk returns a tuple with the PrivateSettlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetPrivateSettlementOk() (*bool, bool) {
	if o == nil || o.PrivateSettlement == nil {
		return nil, false
	}
	return o.PrivateSettlement, true
}

// HasPrivateSettlement returns a boolean if a field has been set.
func (o *CompanyParams) HasPrivateSettlement() bool {
	if o != nil && o.PrivateSettlement != nil {
		return true
	}

	return false
}

// SetPrivateSettlement gets a reference to the given bool and assigns it to the PrivateSettlement field.
func (o *CompanyParams) SetPrivateSettlement(v bool) {
	o.PrivateSettlement = &v
}

// GetPhone1 returns the Phone1 field value if set, zero value otherwise.
func (o *CompanyParams) GetPhone1() string {
	if o == nil || o.Phone1 == nil {
		var ret string
		return ret
	}
	return *o.Phone1
}

// GetPhone1Ok returns a tuple with the Phone1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetPhone1Ok() (*string, bool) {
	if o == nil || o.Phone1 == nil {
		return nil, false
	}
	return o.Phone1, true
}

// HasPhone1 returns a boolean if a field has been set.
func (o *CompanyParams) HasPhone1() bool {
	if o != nil && o.Phone1 != nil {
		return true
	}

	return false
}

// SetPhone1 gets a reference to the given string and assigns it to the Phone1 field.
func (o *CompanyParams) SetPhone1(v string) {
	o.Phone1 = &v
}

// GetPhone2 returns the Phone2 field value if set, zero value otherwise.
func (o *CompanyParams) GetPhone2() string {
	if o == nil || o.Phone2 == nil {
		var ret string
		return ret
	}
	return *o.Phone2
}

// GetPhone2Ok returns a tuple with the Phone2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetPhone2Ok() (*string, bool) {
	if o == nil || o.Phone2 == nil {
		return nil, false
	}
	return o.Phone2, true
}

// HasPhone2 returns a boolean if a field has been set.
func (o *CompanyParams) HasPhone2() bool {
	if o != nil && o.Phone2 != nil {
		return true
	}

	return false
}

// SetPhone2 gets a reference to the given string and assigns it to the Phone2 field.
func (o *CompanyParams) SetPhone2(v string) {
	o.Phone2 = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *CompanyParams) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetFaxOk() (*string, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *CompanyParams) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *CompanyParams) SetFax(v string) {
	o.Fax = &v
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *CompanyParams) GetZipcode() string {
	if o == nil || o.Zipcode == nil {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetZipcodeOk() (*string, bool) {
	if o == nil || o.Zipcode == nil {
		return nil, false
	}
	return o.Zipcode, true
}

// HasZipcode returns a boolean if a field has been set.
func (o *CompanyParams) HasZipcode() bool {
	if o != nil && o.Zipcode != nil {
		return true
	}

	return false
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *CompanyParams) SetZipcode(v string) {
	o.Zipcode = &v
}

// GetPrefectureCode returns the PrefectureCode field value if set, zero value otherwise.
func (o *CompanyParams) GetPrefectureCode() int32 {
	if o == nil || o.PrefectureCode == nil {
		var ret int32
		return ret
	}
	return *o.PrefectureCode
}

// GetPrefectureCodeOk returns a tuple with the PrefectureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetPrefectureCodeOk() (*int32, bool) {
	if o == nil || o.PrefectureCode == nil {
		return nil, false
	}
	return o.PrefectureCode, true
}

// HasPrefectureCode returns a boolean if a field has been set.
func (o *CompanyParams) HasPrefectureCode() bool {
	if o != nil && o.PrefectureCode != nil {
		return true
	}

	return false
}

// SetPrefectureCode gets a reference to the given int32 and assigns it to the PrefectureCode field.
func (o *CompanyParams) SetPrefectureCode(v int32) {
	o.PrefectureCode = &v
}

// GetStreetName1 returns the StreetName1 field value if set, zero value otherwise.
func (o *CompanyParams) GetStreetName1() string {
	if o == nil || o.StreetName1 == nil {
		var ret string
		return ret
	}
	return *o.StreetName1
}

// GetStreetName1Ok returns a tuple with the StreetName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetStreetName1Ok() (*string, bool) {
	if o == nil || o.StreetName1 == nil {
		return nil, false
	}
	return o.StreetName1, true
}

// HasStreetName1 returns a boolean if a field has been set.
func (o *CompanyParams) HasStreetName1() bool {
	if o != nil && o.StreetName1 != nil {
		return true
	}

	return false
}

// SetStreetName1 gets a reference to the given string and assigns it to the StreetName1 field.
func (o *CompanyParams) SetStreetName1(v string) {
	o.StreetName1 = &v
}

// GetStreetName2 returns the StreetName2 field value if set, zero value otherwise.
func (o *CompanyParams) GetStreetName2() string {
	if o == nil || o.StreetName2 == nil {
		var ret string
		return ret
	}
	return *o.StreetName2
}

// GetStreetName2Ok returns a tuple with the StreetName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetStreetName2Ok() (*string, bool) {
	if o == nil || o.StreetName2 == nil {
		return nil, false
	}
	return o.StreetName2, true
}

// HasStreetName2 returns a boolean if a field has been set.
func (o *CompanyParams) HasStreetName2() bool {
	if o != nil && o.StreetName2 != nil {
		return true
	}

	return false
}

// SetStreetName2 gets a reference to the given string and assigns it to the StreetName2 field.
func (o *CompanyParams) SetStreetName2(v string) {
	o.StreetName2 = &v
}

// GetInvoiceLayout returns the InvoiceLayout field value if set, zero value otherwise.
func (o *CompanyParams) GetInvoiceLayout() string {
	if o == nil || o.InvoiceLayout == nil {
		var ret string
		return ret
	}
	return *o.InvoiceLayout
}

// GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetInvoiceLayoutOk() (*string, bool) {
	if o == nil || o.InvoiceLayout == nil {
		return nil, false
	}
	return o.InvoiceLayout, true
}

// HasInvoiceLayout returns a boolean if a field has been set.
func (o *CompanyParams) HasInvoiceLayout() bool {
	if o != nil && o.InvoiceLayout != nil {
		return true
	}

	return false
}

// SetInvoiceLayout gets a reference to the given string and assigns it to the InvoiceLayout field.
func (o *CompanyParams) SetInvoiceLayout(v string) {
	o.InvoiceLayout = &v
}

// GetAmountFraction returns the AmountFraction field value if set, zero value otherwise.
func (o *CompanyParams) GetAmountFraction() int32 {
	if o == nil || o.AmountFraction == nil {
		var ret int32
		return ret
	}
	return *o.AmountFraction
}

// GetAmountFractionOk returns a tuple with the AmountFraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetAmountFractionOk() (*int32, bool) {
	if o == nil || o.AmountFraction == nil {
		return nil, false
	}
	return o.AmountFraction, true
}

// HasAmountFraction returns a boolean if a field has been set.
func (o *CompanyParams) HasAmountFraction() bool {
	if o != nil && o.AmountFraction != nil {
		return true
	}

	return false
}

// SetAmountFraction gets a reference to the given int32 and assigns it to the AmountFraction field.
func (o *CompanyParams) SetAmountFraction(v int32) {
	o.AmountFraction = &v
}

// GetIndustryClass returns the IndustryClass field value if set, zero value otherwise.
func (o *CompanyParams) GetIndustryClass() string {
	if o == nil || o.IndustryClass == nil {
		var ret string
		return ret
	}
	return *o.IndustryClass
}

// GetIndustryClassOk returns a tuple with the IndustryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetIndustryClassOk() (*string, bool) {
	if o == nil || o.IndustryClass == nil {
		return nil, false
	}
	return o.IndustryClass, true
}

// HasIndustryClass returns a boolean if a field has been set.
func (o *CompanyParams) HasIndustryClass() bool {
	if o != nil && o.IndustryClass != nil {
		return true
	}

	return false
}

// SetIndustryClass gets a reference to the given string and assigns it to the IndustryClass field.
func (o *CompanyParams) SetIndustryClass(v string) {
	o.IndustryClass = &v
}

// GetIndustryCode returns the IndustryCode field value if set, zero value otherwise.
func (o *CompanyParams) GetIndustryCode() string {
	if o == nil || o.IndustryCode == nil {
		var ret string
		return ret
	}
	return *o.IndustryCode
}

// GetIndustryCodeOk returns a tuple with the IndustryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetIndustryCodeOk() (*string, bool) {
	if o == nil || o.IndustryCode == nil {
		return nil, false
	}
	return o.IndustryCode, true
}

// HasIndustryCode returns a boolean if a field has been set.
func (o *CompanyParams) HasIndustryCode() bool {
	if o != nil && o.IndustryCode != nil {
		return true
	}

	return false
}

// SetIndustryCode gets a reference to the given string and assigns it to the IndustryCode field.
func (o *CompanyParams) SetIndustryCode(v string) {
	o.IndustryCode = &v
}

// GetDefaultWalletAccountId returns the DefaultWalletAccountId field value if set, zero value otherwise.
func (o *CompanyParams) GetDefaultWalletAccountId() int32 {
	if o == nil || o.DefaultWalletAccountId == nil {
		var ret int32
		return ret
	}
	return *o.DefaultWalletAccountId
}

// GetDefaultWalletAccountIdOk returns a tuple with the DefaultWalletAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetDefaultWalletAccountIdOk() (*int32, bool) {
	if o == nil || o.DefaultWalletAccountId == nil {
		return nil, false
	}
	return o.DefaultWalletAccountId, true
}

// HasDefaultWalletAccountId returns a boolean if a field has been set.
func (o *CompanyParams) HasDefaultWalletAccountId() bool {
	if o != nil && o.DefaultWalletAccountId != nil {
		return true
	}

	return false
}

// SetDefaultWalletAccountId gets a reference to the given int32 and assigns it to the DefaultWalletAccountId field.
func (o *CompanyParams) SetDefaultWalletAccountId(v int32) {
	o.DefaultWalletAccountId = &v
}

// GetFiscalYears returns the FiscalYears field value if set, zero value otherwise.
func (o *CompanyParams) GetFiscalYears() CompanyParamsFiscalYears {
	if o == nil || o.FiscalYears == nil {
		var ret CompanyParamsFiscalYears
		return ret
	}
	return *o.FiscalYears
}

// GetFiscalYearsOk returns a tuple with the FiscalYears field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyParams) GetFiscalYearsOk() (*CompanyParamsFiscalYears, bool) {
	if o == nil || o.FiscalYears == nil {
		return nil, false
	}
	return o.FiscalYears, true
}

// HasFiscalYears returns a boolean if a field has been set.
func (o *CompanyParams) HasFiscalYears() bool {
	if o != nil && o.FiscalYears != nil {
		return true
	}

	return false
}

// SetFiscalYears gets a reference to the given CompanyParamsFiscalYears and assigns it to the FiscalYears field.
func (o *CompanyParams) SetFiscalYears(v CompanyParamsFiscalYears) {
	o.FiscalYears = &v
}

func (o CompanyParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameKana != nil {
		toSerialize["name_kana"] = o.NameKana
	}
	if o.ContactName != nil {
		toSerialize["contact_name"] = o.ContactName
	}
	if o.HeadCount != nil {
		toSerialize["head_count"] = o.HeadCount
	}
	if o.CorporateNumber != nil {
		toSerialize["corporate_number"] = o.CorporateNumber
	}
	if o.TxnNumberFormat != nil {
		toSerialize["txn_number_format"] = o.TxnNumberFormat
	}
	if o.PrivateSettlement != nil {
		toSerialize["private_settlement"] = o.PrivateSettlement
	}
	if o.Phone1 != nil {
		toSerialize["phone1"] = o.Phone1
	}
	if o.Phone2 != nil {
		toSerialize["phone2"] = o.Phone2
	}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.Zipcode != nil {
		toSerialize["zipcode"] = o.Zipcode
	}
	if o.PrefectureCode != nil {
		toSerialize["prefecture_code"] = o.PrefectureCode
	}
	if o.StreetName1 != nil {
		toSerialize["street_name1"] = o.StreetName1
	}
	if o.StreetName2 != nil {
		toSerialize["street_name2"] = o.StreetName2
	}
	if o.InvoiceLayout != nil {
		toSerialize["invoice_layout"] = o.InvoiceLayout
	}
	if o.AmountFraction != nil {
		toSerialize["amount_fraction"] = o.AmountFraction
	}
	if o.IndustryClass != nil {
		toSerialize["industry_class"] = o.IndustryClass
	}
	if o.IndustryCode != nil {
		toSerialize["industry_code"] = o.IndustryCode
	}
	if o.DefaultWalletAccountId != nil {
		toSerialize["default_wallet_account_id"] = o.DefaultWalletAccountId
	}
	if o.FiscalYears != nil {
		toSerialize["fiscal_years"] = o.FiscalYears
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyParams struct {
	value *CompanyParams
	isSet bool
}

func (v NullableCompanyParams) Get() *CompanyParams {
	return v.value
}

func (v *NullableCompanyParams) Set(val *CompanyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyParams(val *CompanyParams) *NullableCompanyParams {
	return &NullableCompanyParams{value: val, isSet: true}
}

func (v NullableCompanyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


