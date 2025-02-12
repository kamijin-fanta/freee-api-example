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

// PartnerCreateParamsAddressAttributes struct for PartnerCreateParamsAddressAttributes
type PartnerCreateParamsAddressAttributes struct {
	// 郵便番号（8文字以内）
	Zipcode *string `json:"zipcode,omitempty"`
	// 都道府県コード（0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄
	PrefectureCode *int32 `json:"prefecture_code,omitempty"`
	// 市区町村・番地（255文字以内）
	StreetName1 *string `json:"street_name1,omitempty"`
	// 建物名・部屋番号など（255文字以内）
	StreetName2 *string `json:"street_name2,omitempty"`
}

// NewPartnerCreateParamsAddressAttributes instantiates a new PartnerCreateParamsAddressAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCreateParamsAddressAttributes() *PartnerCreateParamsAddressAttributes {
	this := PartnerCreateParamsAddressAttributes{}
	return &this
}

// NewPartnerCreateParamsAddressAttributesWithDefaults instantiates a new PartnerCreateParamsAddressAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCreateParamsAddressAttributesWithDefaults() *PartnerCreateParamsAddressAttributes {
	this := PartnerCreateParamsAddressAttributes{}
	return &this
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *PartnerCreateParamsAddressAttributes) GetZipcode() string {
	if o == nil || o.Zipcode == nil {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParamsAddressAttributes) GetZipcodeOk() (*string, bool) {
	if o == nil || o.Zipcode == nil {
		return nil, false
	}
	return o.Zipcode, true
}

// HasZipcode returns a boolean if a field has been set.
func (o *PartnerCreateParamsAddressAttributes) HasZipcode() bool {
	if o != nil && o.Zipcode != nil {
		return true
	}

	return false
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *PartnerCreateParamsAddressAttributes) SetZipcode(v string) {
	o.Zipcode = &v
}

// GetPrefectureCode returns the PrefectureCode field value if set, zero value otherwise.
func (o *PartnerCreateParamsAddressAttributes) GetPrefectureCode() int32 {
	if o == nil || o.PrefectureCode == nil {
		var ret int32
		return ret
	}
	return *o.PrefectureCode
}

// GetPrefectureCodeOk returns a tuple with the PrefectureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParamsAddressAttributes) GetPrefectureCodeOk() (*int32, bool) {
	if o == nil || o.PrefectureCode == nil {
		return nil, false
	}
	return o.PrefectureCode, true
}

// HasPrefectureCode returns a boolean if a field has been set.
func (o *PartnerCreateParamsAddressAttributes) HasPrefectureCode() bool {
	if o != nil && o.PrefectureCode != nil {
		return true
	}

	return false
}

// SetPrefectureCode gets a reference to the given int32 and assigns it to the PrefectureCode field.
func (o *PartnerCreateParamsAddressAttributes) SetPrefectureCode(v int32) {
	o.PrefectureCode = &v
}

// GetStreetName1 returns the StreetName1 field value if set, zero value otherwise.
func (o *PartnerCreateParamsAddressAttributes) GetStreetName1() string {
	if o == nil || o.StreetName1 == nil {
		var ret string
		return ret
	}
	return *o.StreetName1
}

// GetStreetName1Ok returns a tuple with the StreetName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParamsAddressAttributes) GetStreetName1Ok() (*string, bool) {
	if o == nil || o.StreetName1 == nil {
		return nil, false
	}
	return o.StreetName1, true
}

// HasStreetName1 returns a boolean if a field has been set.
func (o *PartnerCreateParamsAddressAttributes) HasStreetName1() bool {
	if o != nil && o.StreetName1 != nil {
		return true
	}

	return false
}

// SetStreetName1 gets a reference to the given string and assigns it to the StreetName1 field.
func (o *PartnerCreateParamsAddressAttributes) SetStreetName1(v string) {
	o.StreetName1 = &v
}

// GetStreetName2 returns the StreetName2 field value if set, zero value otherwise.
func (o *PartnerCreateParamsAddressAttributes) GetStreetName2() string {
	if o == nil || o.StreetName2 == nil {
		var ret string
		return ret
	}
	return *o.StreetName2
}

// GetStreetName2Ok returns a tuple with the StreetName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCreateParamsAddressAttributes) GetStreetName2Ok() (*string, bool) {
	if o == nil || o.StreetName2 == nil {
		return nil, false
	}
	return o.StreetName2, true
}

// HasStreetName2 returns a boolean if a field has been set.
func (o *PartnerCreateParamsAddressAttributes) HasStreetName2() bool {
	if o != nil && o.StreetName2 != nil {
		return true
	}

	return false
}

// SetStreetName2 gets a reference to the given string and assigns it to the StreetName2 field.
func (o *PartnerCreateParamsAddressAttributes) SetStreetName2(v string) {
	o.StreetName2 = &v
}

func (o PartnerCreateParamsAddressAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullablePartnerCreateParamsAddressAttributes struct {
	value *PartnerCreateParamsAddressAttributes
	isSet bool
}

func (v NullablePartnerCreateParamsAddressAttributes) Get() *PartnerCreateParamsAddressAttributes {
	return v.value
}

func (v *NullablePartnerCreateParamsAddressAttributes) Set(val *PartnerCreateParamsAddressAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCreateParamsAddressAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCreateParamsAddressAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCreateParamsAddressAttributes(val *PartnerCreateParamsAddressAttributes) *NullablePartnerCreateParamsAddressAttributes {
	return &NullablePartnerCreateParamsAddressAttributes{value: val, isSet: true}
}

func (v NullablePartnerCreateParamsAddressAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCreateParamsAddressAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


