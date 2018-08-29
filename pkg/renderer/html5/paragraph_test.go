package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("paragraphs", func() {

	Context("regular paragraphs", func() {

		It("a standalone paragraph with special character", func() {
			actualContent := `*bold content* 
& more content afterwards`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong>
&amp; more content afterwards</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a standalone paragraph with trailing spaces", func() {
			actualContent := `*bold content*    
   & more content afterwards...`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong>
   &amp; more content afterwards&#8230;&#8203;</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("a standalone paragraph with an ID and a title", func() {
			actualContent := `[#foo]
.a title
*bold content* with more content afterwards...`
			expectedResult := `<div id="foo" class="paragraph">
<div class="doctitle">a title</div>
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("2 paragraphs and blank line", func() {
			actualContent := `
*bold content* with more content afterwards...

and here another paragraph

`
			expectedResult := `<div class="paragraph">
<p><strong>bold content</strong> with more content afterwards&#8230;&#8203;</p>
</div>
<div class="paragraph">
<p>and here another paragraph</p>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition paragraphs", func() {
		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
this is a note.
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline warning admonition paragraph", func() {
			actualContent := `WARNING: this is a multiline
warning!`
			expectedResult := `<div class="admonitionblock warning">
<table>
<tr>
<td class="icon">
<div class="title">Warning</div>
</td>
<td class="content">
this is a multiline
warning!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("admonition note paragraph with id and title", func() {
			actualContent := `[[foo]]
.bar
NOTE: this is a note.`
			expectedResult := `<div id="foo" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a note.
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("admonition paragraphs", func() {
		It("simple caution admonition paragraph", func() {
			actualContent := `[CAUTION] 
this is a caution!`
			expectedResult := `<div class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
this is a caution!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiline caution admonition paragraph with title and id", func() {
			actualContent := `[[foo]]
[CAUTION] 
.bar
this is a
*caution*!`
			expectedResult := `<div id="foo" class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
<div class="title">bar</div>
this is a
<strong>caution</strong>!
</td>
</tr>
</table>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
