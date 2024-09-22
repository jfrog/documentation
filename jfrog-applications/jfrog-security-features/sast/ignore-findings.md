# Ignore Findings

#### Ignore a Specific Finding

The SAST scanner allows you to ignore a vulnerability finding simply by placing an `jfrog-ignore` annotation directly in the code. Place the `jfrog-ignore` annotation as a comment above the `sink` line of the vulnerability (the final line in the data flow).

The following example shows how to ignore an unsafe-deserialization issue:

<pre class="language-javascript" data-line-numbers><code class="lang-javascript"><a data-footnote-ref href="#user-content-fn-1">export: (req, res)</a> => {
    res = set_cors(req, res)
    res.set('Cache-Control', 'no-store, no-cache, must-revalidate, private');
    payload = Buffer.from(req.body.data, "base64");
    <a data-footnote-ref href="#user-content-fn-2">// jfrog-ignore</a>
    var data = <a data-footnote-ref href="#user-content-fn-3">serialize.unserialize(payload.toString())</a>;
</code></pre>

> **Bottom line:** Place the `jfrog-ignore` annotation above the finding's final line

#### Unignore Findings

To un-ignore a scan finding, simply remove the `jfrog-ignore` annotation from the finding's execution line and re-scan the project.

[^1]: The vulnerability starts here

[^2]: Adding ignore annotation to ignore the finding

[^3]: The vulnerability sink
