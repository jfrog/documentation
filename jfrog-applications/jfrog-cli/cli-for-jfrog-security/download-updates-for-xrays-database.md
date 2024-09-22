# Download Updates for Xray's Database

The offline-update command downloads updates to Xray's vulnerabilities database. The Xray UI allows building the command structure for you.

<table>
   <thead>
      <tr>
         <th width="314.5"></th>
         <th></th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>Command name</td>
         <td>xr offline-update</td>
      </tr>
      <tr>
         <td>Abbreviation</td>
         <td>xr ou</td>
      </tr>
      <tr>
         <td>Command options</td>
         <td></td>
      </tr>
      <tr>
         <td>--license-id</td>
         <td>[Mandatory]<br>Xray license ID.</td>
      </tr>
      <tr>
         <td>--from</td>
         <td>[Optional]<br>From update date in YYYY-MM-DD format.</td>
      </tr>
      <tr>
         <td>--to</td>
         <td>[Optional]<br>To update date in YYYY-MM-DD format.</td>
      </tr>
      <tr>
         <td>--version</td>
         <td>[Optional]<br>Xray API version.</td>
      </tr>
      <tr>
         <td>--target</td>
         <td>[Default: ./]<br>Path for downloaded update files.</td>
      </tr>
      <tr>
         <td>--stream</td>
         <td>[Default: false]<br>Set to true to use Xray DBSync V3 stream, Possible values are: public_data, exposures and contextual_analysis.</td>
      </tr>
      <tr>
         <td>--periodic</td>
         <td>[Default: false]<br>Set to true to get the Xray DBSync V3 Periodic Package (Use with stream flag).</td>
      </tr>
      <tr>
         <td>Command arguments</td>
         <td>The command accepts no arguments.</td>
      </tr>
   </tbody>
</table>