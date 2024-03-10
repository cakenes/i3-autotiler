### Autotiler for i3wm https://i3wm.org/

Simple autotiler written in <b>go</b> & <b>bash</b>, pick your poison.<br />
<code>autotiler.service</code> looks for <b>go</b> version by default, but I've included <b>bash</b> version as well ( modify service & copy over .sh instead .)<br />
Will listen for events from <code>i3-msg</code> and pick split direction based on focus window height & width.<br />

### Installing:
<code>i3-autotiler.service</code> should be moved to: <code>/etc/systemd/user/</code><br />
<code>i3-autotiler</code> should be moved to <code>/usr/local/bin/</code><br />

<code>systemctl daemon-reload</code><br />
<code>systemctl enable --user --now i3-autotiler.service</code> to start.
