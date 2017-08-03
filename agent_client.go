package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var fs *flag.FlagSet

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

//rand.Seed(42) // Try changing this number!
var messages []string

func main() {

	// Flag
	fs = flag.NewFlagSet("", flag.ExitOnError)
	var (
		malformed = fs.Int("malformed", 0, "0: Formed, 1: Malformed")
		count     = fs.Int("count", 1, "Count to send")
	)
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	//	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:513")
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:19902")
	CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	// Connect to server
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	defer Conn.Close()
	//	msg := "<34>1 2017-02-03T22:14:15.003Z hello.com application - MSGID47 Message from UDP client   "

	//
	initMessages()
	rand.Seed(100)

	i := 1
	if *malformed == 0 {

		//		for i <= *count {
		//			msg := fmt.Sprintf("<%d>%d %s %s %s %d %s %s",
		//				random(0, 23), // priority
		//				1,             // version
		//				time.Now().Format("2006-01-02T15:04:05Z"), // timestamp
		//				"udpclient.com",                           // hostname
		//				"udp_client",                              // app-name
		//				random(1000, 9999),                        // process id
		//				"ID10",                                    // message id
		//				getRandMessage(),
		//			)
		//			_, err = Conn.Write([]byte(msg))
		//			fmt.Println(msg)
		//			i++
		//		}
		src := []byte("5f3d4526d15a37cf8243103b6004b3a13ff8abe735ecc788d4879f3bef34a92ce446cb97aed9350704351b27dfb7e851991ad101b0be39154165c61856be2f178513d057024eb8b628dfca77607742d68206c20667b6a54fb467bdbbd2df71ab1e4430bf4ad279db3d08332c55d12f05e1e996a46d11d9c753f845eb87b1c1189f0b3af3057c9dd657fbde1ac637cf62")
		dst := make([]byte, hex.DecodedLen(len(src)))
		n, _ := hex.Decode(dst, src)

		//		aa := []byte("5f3d4526d15a37cf8243103b6004b3a13ff8abe735ecc788d4879f3bef34a92ce446cb97aed9350704351b27dfb7e851991ad101b0be39154165c61856be2f178513d057024eb8b628dfca77607742d68206c20667b6a54fb467bdbbd2df71ab1e4430bf4ad279db3d08332c55d12f05e1e996a46d11d9c753f845eb87b1c1189f0b3af3057c9dd657fbde1ac637cf62")
		_, err = Conn.Write(dst[:n])

		//		_, err = Conn.Write([]byte(`<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8`))
		//		_, err = Conn.Write([]byte(`<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.`))
		//		_, err = Conn.Write([]byte(`<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] BOMAn application event log entry...`))
		//		_, err = Conn.Write([]byte(`<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"][examplePriority@32473class="high"]`))

	} else {
		for i <= *count {
			_, err = Conn.Write([]byte("Malformed log"))
			i++
		}
	}

	//	if err != nil {
	//		fmt.Println(msg, err)
	//	}
}

func printHelp() {
	fmt.Println("udpclient [options]")
	fs.PrintDefaults()
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func getRandMessage() string {
	return messages[rand.Intn(len(messages))]
}

func initMessages() {
	messages = []string{
		"- ADD: ADMINISTRATION: New Management Console with significant functionality, speed and usability improvements",
		"- ADD: ANTISPAM: Integrated SpamAssasin Module allows you to call SpamAssasin Daemon and integrate results into Spam/Message Filters",
		"- ADD: ANTISPAM: PTR Filtering can now be used in system message filtering (it used be available as at the SMTP connector for blocking)",
		"- ADD: ANTISPAM: Report As Spam option in webmail can now signal bayesian engine to process Spam messages",
		"- ADD: CALDAV: CalDAV is now integrated as a service under IIS (so it no longer requires separate IP address). It is serviceable through MailEnable Protocols Site",
		"- ADD: CARDDAV: CardDAV is now integrated as a service under IIS (so it no longer requires separate IP address). It is serviceable through MailEnable Protocols Site",
		"- ADD: MIGRATION: Generic IMAP Migration Agent to migrate from any IMAP/POP service",
		"- ADD: MIGRATION: Migration Agent supplied allowing migration from IceWarp",
		"- ADD: MIGRATION: Migration Agent supplied allowing migration from Ipswitch IMail",
		"- ADD: MIGRATION: Migration Agent supplied allowing migration from Microsoft Exchange",
		"- ADD: MIGRATION: Migration Console allows you to simply migrate from other message systems, with tracked progress and reporting.",
		"- ADD: MOBILE: New Mobile WebMail client (JQuery Mobile) provides a tuned interface for Mobile Devices",
		"- ADD: MONITORING: Tray Monitoring provides detailed Spam Detection/Blocking Statistics",
		"- ADD: PLATFORM: Ability to configure Postoffice level defaults for all new mailboxes",
		"- ADD: PLATFORM: Ability to configure System level defaults for all new postoffices",
		"- ADD: PLATFORM: All services can now run as Native 64 bit services on 64 bit O/S",
		"- ADD: PLATFORM: Can run custom scripts when creating and removing new postoffices/domains",
		"- ADD: PLATFORM: Creation and removal of new postoffices/domains can now provision DNS and IIS automatically",
		"- ADD: PLATFORM: New Monitoring Dashboard within MMC (displays messages for critical system events and alerts)",
		"- ADD: PLATFORM: Registry Setting provided to throttle IMAP connections if they attempt to overload the server",
		"- ADD: SMTP: Developers can now drop e-mail messages in a designated folder and the system will route them to their destination",
		"- ADD: SYNCML: SyncML is now integrated as a service under IIS (so it no longer requires separate IP address). It is serviceable through MailEnable Protocols Site",
		"- ADD: WEBDAV: MyFiles is now integrated as a service under IIS (so it no longer requires separate IP address). It is serviceable through MailEnable Protocols Site",
		"- ADD: WEBMAIL: New Skins Library allows you to download and apply skins from MailEnable's Online Skin Library",
		"- ADD: WEBMAIL: Standard Edition now includes Calendaring Support within its webmail client",
		"- ADD: WEBMAIL: Webmail inbox will now show SMS messages that have been synchronised from device via EAS or other protocol",
		"- IMP: INSTALL: Installation no longer requires the pre-requisite of IIS Legacy Scripting libraries",
		"- IMP: WEBADMIN: Improved layout of WebAdmin",
		"- IMP: WEBMAIL: WebMail enhancements (including Notes, Sender Identities, Multiple File Download and List-Unsubscribe)",
		"ADD: Added ability to disable Basic authentication for Synchronisation Protocols",
		"ADD: Added Block IP to the MMC Admin queued Message Details screen (so you can block the client IP address from connecting to the SMTP service)",
		"ADD: Added client IP, authenticating user, and authentication status in the queued Message Details window within the MMC Admin.",
		"ADD: Added Digest authentication for all Synchronisation Protocols (specifically for CalDAV/CardDAV since Apple devices with current OSX releases require it)",
		"ADD: Added Disable Mailbox from the queued Message Details screen (an administrator can disable a mailbox directly from the Message Details screen)",
		"ADD: Added option under Messaging Manager Properties/Administration to allow (or disallow) asynchronous population of lists within the MMC)",
		"ADD: Added View Send History (so an administrator can view the send history from within the queued Message Details screen)",
		"ADD: CardDAV now publishes Global Address List (compatible with OpenProtocols Connector)",
		"ADD: Details for queue items in admin has more details and ability to block mailbox authenticating",
		"ADD: Improvements to LDAP address list report (behaviour can be configured within Management Console)",
		"ADD: Message Tracker is available through the admin program under the Diagnose branch",
		"ADD: Microsoft ActiveSync will now allow connections from Microsoft Outlook 2013",
		"ADD: Support for XLIST for IMAP service",
		"ADD: SYSTEM: Performance counters now available for 64 bit Performance Monitor",
		"EAS: EAS  - BB10 would not immediately recognise new message arrivals, since it issues folder polls and Sync requests concurrently. ",
		"FIX Postoffice IP Bindings may not initally show correct status if it had not been configured (with new MMC3)",
		"FIX: (EAS) Editing a task on the Server would not always update on client devices",
		"FIX: 64bit version of MEAVGEN.DLL was not installed on 64 bit servers",
		"FIX: A delay of around 30 seconds may occur when attempting to use Postoffice default settings (when provisioning is enabled)",
		"FIX: Accepting meeting change in webmail was not updating the VCAL times in the underlying appointment data (VCALENDAR)",
		"FIX: Accessing GAL via web mail or adding a Directory member in admin would produce a .NET error.",
		"FIX: Accessing personal contacts via LDAP may cause server exception/crash",
		"FIX: Active Sync may duplicating Emails in the Deleted Items folder with Outlook 2013",
		"FIX: ActiveSync (EAS) Sync was returning empty sync response if the same sync key is used (but different request). It should only return an empty response if the request data has not changed.",
		"FIX: ActiveSync could send item change notifications to items deleted in the same request (since BB10 issues change and deletes in the same request)",
		"FIX: ActiveSync did not add friendly name to From header if it is missing when attempting to send",
		"FIX: ActiveSync may not clear all device settings/state configuration when device is reconfigured (potentially causing duplicates in Microsoft Outlook/EAS)",
		"FIX: ActiveSync message replies or forwards sent from apple that include a mixed array of recipient formats (Aliases, anglebraced and comma delimited literals) may not be delivered to all recipients.",
		"FIX: ActiveSync much faster processing sent emails",
		"FIX: ActiveSync will now attempt to throttle the maximum response size to 20MB",
		"FIX: ActiveSync will now permit previous sync key re-use when no changes have been detected (reducing the need to resync folders upon communication error)",
		"FIX: ActiveSync will now use incremental SyncKeys to protect agains communication failure when processing MoreAvailable (typically occurs when initially loading large mailboxes with Outlook)",
		"FIX: ActiveSync Windows Phone may display synchronisation error upon account creation",
		"FIX: ActiveSync would not encode 7bit attachments in base64 format (which is a requirement of EAS)",
		"FIX: ActiveSync would not use Mailbox TimeZone setting as the default when viewing Free and Busy times of All Day Events",
		"FIX: Added additional Hotmail specific EAS schema extensions to allow the server coupling of the Junk Mail folder",
		"FIX: Adding a new item under MyFiles would not immediately show in DAV clients",
		"FIX: Adding extended characters to Auto Response in web mail did not encode characters correctly in MMC",
		"FIX: Address Resolution Problem after uploading multiple attachents in web mail",
		"FIX: Admin not showing sender or recipient address in message details in queues if they contained = symbol",
		"FIX: Admin outbound mail view send history option was not reading logs that were currently active by the mail services",
		"FIX: Admin program would generate error when testing an invalid list datasource",
		"FIX: Admin was showing the SMTP welcome message when viewing POP options",
		"FIX: Advanced options settings page under webmail would not be visibile if ActiveSync was disabled for postoffice",
		"FIX: An apostrophe in the email address would stop the read page loading",
		"FIX: Apostrophe in email address in message would prevent email contents showing in webmail when viewing email in new window",
		"FIX: Appointments generated via old Firefox/Lightning clients may not be able to be parsed and could corrupt extended characters within data fields",
		"FIX: Auto-complete not working in standard Edition Webmail",
		"FIX: Autocomplete fails to resolve when invalid email address exists within a contact",
		"FIX: Autoresponder now logs to the postoffice connector debug log that it fired",
		"FIX: Autoresponders configured via MAPI may not include sender address if the default sender address was not defined for the mailbox",
		"FIX: Base64 encoded calendar invites that were attachments were not detected as invites when sent to EAS clients",
		"FIX: Blank E-mail messages with only an attachment would not save to Drafts Folder",
		"FIX: BlockFileLock Activity notification handler (WaitForAccessToFile) failed to obtain a lock on the blockfile may be misreported in Windows Event Log",
		"FIX: Bug fixes to vcard and vcalendar encoding (conversions from version 2.1 to version 3)",
		"FIX: CalDAV ampersands (and other characters) rendered in xml were being escaped even though the data was in a CDATA section",
		"FIX: CalDAV would fail to send result if the calendar file on server did not have a VCALENDAR",
		"FIX: Can now manage EAS server cache from within WebMail (same as with SyncML cache)",
		"FIX: Cannot open messages in new window in webmail if the folder name has a quote in it",
		"FIX: Cant a new postoffice when using MySQL as repository in 7.50.",
		"FIX: CardDAV now handles contacts containing legacy (@) character",
		"FIX: CardDAV on OSX would not process all vCards",
		"FIX: Changed the ordering of EAS XML responses since WP8 seems to require specific ordering (any changes that occur while synchronising may cause sync error)",
		"FIX: Changes to folder contents may unnecessarily set the a change flag file, potentially causing folders to be reindexed and resulting in flag loss and unnecessary replication).",
		"FIX: Clicking Print Preview and View Header buttons may cause action to fire multiple times",
		"FIX: Configuration of XLIST support in IMAP admin did not match existing methods",
		"FIX: Contact photos file content type in webmail is being returned as application/binary",
		"FIX: Contact photos may not render when downloaded via clients requesting VCARD Version 3 (CardDAV on OSX)",
		"FIX: Contacts Photos may be removed when edited via WebMail (depending on the format of the VCARD)",
		"FIX: Core services may leak memory over time",
		"FIX: Critical Security Update for SMTP Connector",
		"FIX: CRITICAL: Security update for MailEnable Messaging Platform (all versions)",
		"FIX: Cross Site Scripting Vulnerability protection on posted forms and AJAX requests",
		"FIX: Default MAI is was not always copied to new mailboxes when created in the new MMC Administration Console",
		"FIX: Deleting a postoffice via the provisioning engine may leave ophaned address entries",
		"FIX: Deleting of Quarantined items in the Management Console would only remove the command file and not the message file",
		"FIX: Deleting of Quarantined items in the Management Console would only remove the command file and not the message file",
		"FIX: Diagnostic utility did not report URL blacklisting status",
		"FIX: Disabling a list in the MMC via the folder tree does not disable it.",
		"FIX: Disabling a mailbox in new MMC admin is not disabling login",
		"FIX: Disabling filter in new mailenable management console would not disable",
		"FIX: Discrete all day events may report as a two days in Android EAS clients ",
		"FIX: DKIM administration was broken in administration program",
		"FIX: DKIM may fail due to the order of header items",
		"FIX: DNS provisioning is failing with exceptions.",
		"FIX: Double clicking filter criteria When message is marked as priority may cause MMC to crash",
		"FIX: Downloading some attachments as a consolidated ZIP may result in a webmail crash",
		"FIX: EAS  - ActiveSync would prematurely respond to empty ping requests - causing battery drain on IOS devices",
		"FIX: EAS calendar issues with Samsung devices. Floating time calendar entries (created on Mac) are not displayed on the device",
		"FIX: EAS License validation may not be able to contact MailEnable's Licensing Server",
		"FIX: EAS meeting events in calendar may not display the correct time if daylight savings change last week of month",
		"FIX: EAS the mailbox name is sent down as the UserDisplayName (It now sends down the actual display name)",
		"FIX: EAS was not sending down the correct TZ recurrence information",
		"FIX: EAS: IOS devices will not show complete message list when they contain items with the SMS message class ",
		"FIX: EAS: Out of Office not implemented for EAS",
		"FIX: Editing a appointment in in web mail in Mountain time shifted the time slots one hour within the appointment editing window",
		"FIX: Editing the alternate port number for IMAP in MMC config did not enable the apply button",
		"FIX: Editor Languages will now work with button label translations",
		"FIX: Empty or corrupt message cache file for administration program could cause an error when viewing system messages",
		"FIX: Entering local email addresses for the domain properties abuse@, etc in MMC admin causes looping",
		"FIX: Error when adding postoffice in admin with provisioning enabled",
		"FIX: Event log error was logged if Outlook Connector failed to log in, to a non-existant mailbox",
		"FIX: Exception may be raised when running web mail under NET 4.0 when webmail accesses a corrupt FOLDER.DAT file",
		"FIX: Exception thrown in MMC when changing webmail default postoffice",
		"FIX: Exchange migration may not be able to find appropriate libraries to facilitate migration",
		"FIX: Exchange Migration would not work reliably with self signed EWS certificates",
		"FIX: Exporting ICS file from webmail using popup on calendar display causes error",
		"FIX: Failed authentications against webmail may raise an exception",
		"FIX: Fixed web administration directory management to support extended character fields",
		"FIX: Folders that contain the name tasks cannot been seen under root folders in IMAP",
		"FIX: Forwarding a message in web mail does not insert the FROM adddress in the original headers if the FROM has a friendly name",
		"FIX: Forwarding a message in webmail would not add Message-Id header when generating compound message",
		"FIX: Group seperators for email headers could prevent the To/From addresses showing in webmail",
		"FIX: Having preview pane at bottom causes multiple toolbar events to fire",
		"FIX: Highlighting links in a plain text email incorrect marks up emails within secure HTML links already marked up",
		"FIX: HTML signatures are now converted to plain text if they were composed as HTML, but the plain text editor is selected.",
		"FIX: HTTPMail Service could crash on some formated calendar items retrieved by CalDAV",
		"FIX: ICS/iCalendar Sharing will not allow access to shared mailbox calendars",
		"FIX: Identities using removed signatures could generate webmail error",
		"FIX: IE 10 may freeze after cancelling a search in WebMail",
		"FIX: IE8 & IE7 compatibility improvements for webmail",
		"FIX: If a client connection sends two emails in the one connection then the server is writing out the first emails subject to both command files.",
		"FIX: IIS Virtual directory configuration of WebMail/WebAdmin would not commit changes to IIS metabase",
		"FIX: IMAP - IMAP Command continuation would not correctly decode credentials when consecutive command continuation requests were issued.",
		"FIX: IMAP 64Bit NTLM driver was not distributed in Premium and Enterprise kits",
		"FIX: IMAP AUTH LOGIN was returning BAD instead of NO for invalid password",
		"FIX: IMAP could return incorrect HasChildren and HasNoChildren flags",
		"FIX: IMAP CRAM-MD5 returned text for failed logins is improved",
		"FIX: IMAP Expunge notifications from webmail may be incorrect ordered (depending on order items are selected within webmail)",
		"FIX: IMAP IDLE after selecting a mailbox with EXAMINE, notifications are not received for the connection",
		"FIX: IMAP may mis-report that the connection limit has been reached and that server is too busy.",
		"FIX: IMAP may not accept connections immediately when service is restarted",
		"FIX: IMAP may raise exceptions when signalling other connections (leading to a restart of the IMAP exception)",
		"FIX: IMAP may raise signalling exception under load on 64 bit servers",
		"FIX: IMAP NOOP command may place additional load on server if folder is marked with changes (via _change.dty file)",
		"FIX: IMAP not longer returns previous folder details on blank SELECT (which may lead to Apple Mail crashing)",
		"FIX: IMAP Postoffice/Domain bindings would not respect mappings",
		"FIX: IMAP SEARCH command not supporting the KEYWORD search key for $Forwarded ",
		"FIX: IMAP Service may experience high CPU and excessive logging when authenticating dropped SSL sessions",
		"FIX: IMAP UID EXPUNGE may not notify other connections with all changes",
		"FIX: IMAP was not listing the clientIP in system manager messages for failed login attempts",
		"FIX: IMAP was returning PERMANENT FLAGS full response for EXAMINE command",
		"FIX: IMAP will now queue notifications for dispatch under idle rather than dispatching them during idle",
		"FIX: IMAP: IMAP may raise an exception if client requests a poorly formatted message range ",
		"FIX: IMAP: IMAP tries to log to the audit log failed logins",
		"FIX: IMAP: Improved handling/scalability of inbound SSL communication with IMAP protocol ",
		"FIX: Improved locking for IMAP subsciption folder list",
		"FIX: Improved migration queue contention to reduce elapse time and impact on active connections",
		"FIX: Improved the HTML parsing of webmail messages (some messages/tags would be stripped when they did not need to be)",
		"FIX: Improved the layout and function of the port bindings for SMTP Connector within the MMC Admin",
		"FIX: Improved the way attach vcs files are attached to emails",
		"FIX: Improvements to EAS with respect to Galaxy S3 (reduces inbox refreshing by formatting change responses differently)",
		"FIX: Improvements to EAS with respect to Galaxy S3 (reduces inbox refreshing by formatting change responses differently)",
		"FIX: Inbound SMTP scripts no longer fail on sender/recipient emails with quotes in them",
		"FIX: INSTALL: MEInstaller is not removing the 3.5 .NET controls from web.config",
		"FIX: Installation on IIS/IIS6 compatability system better detects later versions of the .NET framework",
		"FIX: Installation on Windows 8 or Windows 2012 may fail to create WebMail and WebAdmin web sites",
		"FIX: Installation on Windows 8 or Windows 2012 may fail to create WebMail and WebAdmin web sites",
		"FIX: Installer could fail to configure web admin if root website did not exist",
		"FIX: iPhone/iPad may not list all recipients in To or CC field via EAS",
		"FIX: Large UID COPY IMAP command could cause high CPU",
		"FIX: List owner in admin was not compulsory field",
		"FIX: Long subject or attachment text in webmail was not wrapping, pushing right side of container outside preview pane",
		"FIX: Mailbox default options were not being effectively deployed to new mailboxes",
		"FIX: MailEnable could attempt to create connection reporting files for IMAP,POP and SMTP in root drive (if the registry key was not defined)",
		"FIX: MailEnable Installer may fail attempting to access MAPI32.DLL if it did not exist on target system (Windows Core)",
		"FIX: MailEnable Management Console may raise an exception when attempting to retrieve the CPU counter on some servers",
		"FIX: ME application pools are stopped when you install and upgrade ME on Windows 2012",
		"FIX: MEDiag gives warning for ASPNET account not having permissions on servers without ASPNET account",
		"FIX: Meeting request with timezone offsets comprised with a partial 30 minute offset would ignore the 30 minute offset.",
		"FIX: MEInstaller may crash if Application Pools were not running at the time of executing a configuration action",
		"FIX: MEInstaller was not removing removing attributes properly from web.config file when downgrading from Version 4 to Version 2 of the Framework",
		"FIX: Message details show in administration program had created time and next send time in UTC",
		"FIX: Message forward by iphone (EAS) was not correctly inserting BOM header into original message contents",
		"FIX: Messages would not display in Outlook connecter if they are moved by mailbox filter",
		"FIX: Migration application could produce an error for migrations in progress",
		"FIX: MIGRATION: Migration utility was not setting message flags after migrating messages via IMAP",
		"FIX: Missing option in new MMC (Only generate NDRs for senders who authenticate)",
		"FIX: MMC  - Aligment issues fixed on Mailbox Properties tab",
		"FIX: MMC  - Numeric columns would sort alphabetically rather than numerically",
		"FIX: MMC  - Unencrypting Password storage in Authentication file would not decode properly",
		"FIX: MMC Admin was not creating special folders and defaults for new mailboxes (and relied on services to create some folders)",
		"FIX: MMC error when clicking Groups node",
		"FIX: MMC feature to Forward all mailbox messages to another address would not work correctly in the MMC version.",
		"FIX: MMC: Added generic exception handing in new MMC, that will email exceptions direct to MailEnable",
		"FIX: MMC: Admin Program is missing two action popup menus on quarantine items",
		"FIX: MMC: Could not reliably save custom blacklists in new MMC",
		"FIX: MMC: Deleting from quarantine (and perhaps others) does not remove command and message file ",
		"FIX: MMC: Error message would occur when opening read receipt messages without subjects",
		"FIX: MMC: NET error when creating a new host header for web mail in new mmc when MailEnable Web mail does not exist under IIS",
		"FIX: MMC: No Release button in Quarantine",
		"FIX: MMC: SSL certificate setting was not reliably saving when selected in MMC admin",
		"FIX: MMC: Uninstalling a webmail skin causing MMC exception",
		"FIX: MMC: Viewing a skin in the new MMC admin does not work correctly ",
		"FIX: MMC: With Version 7, all new mail boxes are initially created with a 0 KB quota ",
		"FIX: Mobile webadmin did not add email addresses to a new mailbox",
		"FIX: Mobile webadmin was missing some images",
		"FIX: Mobile webmail back button stops working after time",
		"FIX: Mobile webmail was not allowing multiple recipients",
		"FIX: Mobile webmail would not respect all To/CC/BCC recipients ",
		"FIX: MOBILE: Mobile WebMail Sender Addresss may appear as blank ",
		"FIX: Moving a message that is set with an read flag may not update FOLDER.DAT correctly in destination folder (and may lead to incorrect read count)",
		"FIX: Moving some appointment types in webmail calendar by dragging could lock up webmail",
		"FIX: MTA could crash when global spam protection was doing whitelist check and the command file was an NDR with no IP address in it ",
		"FIX: MTA may raise a null pointer exception if connecting client IP address was not populated (by third party connectors or filter actions)",
		"FIX: multipart/relative is handled same as multipart/related for IMAP BODYSTRUCTURE command",
		"FIX: NETInstaller prompts (in error) to ask to create Mobile as a virtual directory",
		"FIX: New mailboxes created in the MMC were applied with quota flag set",
		"FIX: New mailboxes created in the MMC were applied with quota flag set",
		"FIX: New Management Console,WebAdmin and WebMail now prevent passwords with extended characters",
		"FIX: New Managment Console  may raise an exception if you try and delete a log file that is being used by the service.",
		"FIX: New MMC (MMC3) would not list all IP interfaces - It would only list those associated with the hostname of the main network adapter",
		"FIX: New MMC autotraining dialog for ham/spam addresses always says Ham addresses (even when the Spam dialog is opened)",
		"FIX: New MMC could take a long time to load large numbers of Postoffices (solution is to disable Asynchronous loading)",
		"FIX: New MMC was not able to save save scripted filters on 64 bit installations",
		"FIX: New MMC was not able to save save scripted filters on64 bit installations",
		"FIX: New MMC would not import members for a list from a.csv file",
		"FIX: Not sending up user-agent in webmail login page would raise an exception in WebMail",
		"FIX: Option to automatically change email address when changing a domain name not working",
		"FIX: Outlook EAS 2013 would not report all message header items",
		"FIX: Outlook EAS is would not immediate sync nested subfolders until the client was reopened.",
		"FIX: Photo encoding for CardDAV where VCard boundaries may not be converted property from Version 2.1 to Version 3.0 (and vice versa)",
		"FIX: POP maximum thread thresholds were not being enforced correctly (allowing more active connections than specified)",
		"FIX: POP Non-SSL connection is made to POP on SSL port does not respect session timeout settings",
		"FIX: Popup contact details in webmail for email no longer shows thumbnail",
		"FIX: Possible ActiveSync crash on initial Android sync",
		"FIX: Postoffice connector could fail if the delivery event timed out and had to be forced to terminate",
		"FIX: Postoffice connector option for disabling redirections does not work",
		"FIX: QP encoded descriptions in CalDAV items may not be decoded fully",
		"FIX: Quotes in emails may cause problems with the VBScript in SMTP Filters",
		"FIX: Reclaiming items marked for deletion would leave index.xml file locked, causing handle leak and system lockups",
		"FIX: Right click action for disabling a mailbox in new MMC Admin was not also disabling login",
		"FIX: Samsung IMAP client not showing IMAP folder list",
		"FIX: Samsung S4 (EAS) loses message after moving message to a new sub folder and resyncing",
		"FIX: Security Patch for POP Retrieval Service (Advisory forthcoming)",
		"FIX: Setting a POP retrieval entry to disabled when using MySQL or SQL Server does not disable it unless you remove it.",
		"FIX: setting forward flag in webmail was not updating the index.xml file",
		"FIX: Skins library is now visible by default",
		"FIX: SMTP CRAM-MD5 authentication over SSL may not authenticate on IOS 7 devices",
		"FIX: SMTP EHLO response was not reflecting IP bindings",
		"FIX: SMTP had memory leak in outbound TLS and in logging IPv6 addresses",
		"FIX: SMTP Inbound will truncate recipient list when the Maximum Recipient Threshold is reached (rather than reporting an error to the client)",
		"FIX: SMTP may deliver delay status notifications to some recipients when the message was able to be received (but other recipients failed)",
		"FIX: Some messages could not be blacklisted/whitelisted from message list right click menu",
		"FIX: Some messages may not display correctly in webmail",
		"FIX: Some outbound SMTP queue columns in new admin program were missing",
		"FIX: Sorting lists may cause invalid values to appear against postoffice/mailbox items",
		"FIX: SpamAssassin check in system Spam protection filter did not adjust overall message score",
		"FIX: Synchronization Service now has the alternate port set to be 8443 (rather than showing as zero)",
		"FIX: System Overview notcorrectly reporting the SMTP Outbound Queue Length",
		"FIX: Tasks with notes are not replicated to and from server via ActiveSync",
		"FIX: Tasks with notes are not replicated to and from server via MAPI",
		"FIX: TEL tag in VCF file is not being converted to VCARD 3 format when sent to Yosemite client",
		"FIX: The client configuration page in webmail crashes on some email addresses",
		"FIX: The custom SMTP welcome message is overwriting the POP custom welcome message",
		"FIX: The labels for text boxes on contacts and settings pages were slightly lower than text boxes (in Mobile WebMail)",
		"FIX: The option to check spelling of webmail messages before being sent was not working",
		"FIX: Timezone changes that occur on the last weekday of the month were not being respected",
		"FIX: Toolbar buttons unclickable in Opera, under Contacts & Tasks",
		"FIX: Turkish regional settings on server will remove tags in signatures",
		"FIX: UID expunges are not logged into the audit log",
		"FIX: Umlauts would not show correctly in the mailbox friendly name column in new MMC",
		"FIX: Under load the SMTP service may accept emails when user is over the per hour restriction",
		"FIX: Updated documentation for MEMigrate to refect default EWS path",
		"FIX: Updated invalid SMTP firewall exceptions on initial install",
		"FIX: Updating the status via a scheduling response in WebMail would not always be able to locate the original message (in cases where the Appointment UID was wrapped)",
		"FIX: Upgrading MailEnable may skip the Framework detection on IIS6 but still add the 3.5 dependencies to web.config.",
		"FIX: Uploading contact image does not work on webmail",
		"FIX: Various webmail security issues (thanks to Soroush Dalili from NCC Group for his work on this)",
		"FIX: VCF files directly copied to Contacts folder would not display if the .VCF extension was in lower case",
		"FIX: WebAdmin did not respect default value for maximum mailbox size postoffice setting",
		"FIX: WebAdmin was not creating mailbox special folders upon creation",
		"FIX: WebAdmin welcome page would displays links to some options even if they are disabled",
		"FIX: WEBADMIN: Script error accessing web admin using IE 10 on Windows",
		"FIX: Webmail allowed users to redirect to themselves",
		"FIX: Webmail and mobile webmail now use the SMTP restrictions for number of recipients",
		"FIX: Webmail incorrectly processes recurrence rule for daylight savings changes on last day of a month",
		"FIX: Webmail may generate orphaned file/file locks on Windows 2012 /IIS8 (with 64 bit Application Pools)",
		"FIX: WebMail may produce stack trace if cached message count is out of date",
		"FIX: Webmail upload limit property page in admin would crash if the maximum allowed content length was over 1gb",
		"FIX: Webmail was not showing some attachments",
		"FIX: Webmail would not permit uploading a file with two full stops in it",
		"FIX: WEBMAIL: Deleting selected tasks in web mail could corrupt task index ",
		"FIX: WEBMAIL: Webmail may raise an exception when viewing account summary screen with old Versions of MailEnable base (Hoodoo)",
		"FIX: When blacklist or whitelist editing is disabled for webmail, users can still add entries",
		"FIX: When creating a new postoffice console may freeze for 30 seconds and does not add correct postmaster credentials to AUTH.tab",
		"FIX: When no advert campaign is selected in web admin, advertising banner was showing blank area if advertising enabled",
		"FIX: When reporting a messaage as spam in web mail it does not notify the MAPI client",
		"FIX: When sending a meeting request to a mailbox which is synced via EAS in Outlook 2013, the meeting request is synced as a message rather th a meeting request.",
		"FIX: When Skype click to call plugin is installed, scrolling through webmail contact list does not refresh properly",
		"FIX: When using identities in webmail it was using the mailbox friendly name, not the identity one",
		"FIX: Windows 8 mobile device via EAS may not display a new contact immediately after it has been added.",
		"FIX: Windows phone may not sync when using ActiveSync",
		"FIX: With IE 11, in webmail, autocomplete for email addresses when composing is not working",
		"FIX: Workaround for bug where Blackberry CALDAV clients assume that returned URLs contain ICS extensions (preventing them from listing calendar contents)",
		"FIX: Workaround for Cisco SMTP proxy misbehaviour (to overcome problem sending via SMTP proxy)",
		"FIX; SMTP Delivery Properties - Maximum of 0 connections to the same server (now uses a respectable default value)",
		"IMP: Empty Folder option within webmail does not immediately update mailbox usage",
		"IMP: Access control IP management in admin improvements",
		"IMP: ActiveSync now implements the Preview field when it is requested",
		"IMP: ActiveSync will now handle and retransmit stateful updates if a loss of sync state occurs",
		"IMP: Added additional authentication credential caching to significantly improve the responsiveness of authenticating users",
		"IMP: Added additional caching of store change list to improve EAS performance with Outlook",
		"IMP: An autoresponder now logs to the postoffice connector debug log that it fired",
		"IMP: CalDAV now reports tasks in its own folder (for better integration with iOS and Thunderbird (Thunderbird now requires setting up an additional /task folder for task support)",
		"IMP: Debug log did not record failures when attempting to access CalDAV when it is disabled",
		"IMP: Drop folder is not enabled by default and not configurable within the Management Console",
		"IMP: EAS  - ActiveSync Server will now cap the number of concurrent commands to 5. A Server Busy message will be returned to clients issuing more than 5 concurrent commands",
		"IMP: Exchange Migration via Exchange Web Services will now auto-discover and automatically migrate accounts",
		"IMP: IIS Virtual directory configuration would not grant IME_ADMIN access to parent web.config",
		"IMP: IMAP service now creates special folders (if they do not exist) when a user logs in (so that Outlook 2013 will correctly assign IMAP folders for new mailboxes)",
		"IMP: Improved handing of Ping Collisions to reduce bandwidth when using ActiveSync with Outlook 2013",
		"IMP: Improved installation and prerequisite checking for Windows 2008 and Windows 2012 installations",
		"IMP: Improved logging for protocols managed under IIS (ActiveSync, calDAV, cardDAV, etc) - there are additional logs under the HTTPMail logging directory",
		"IMP: Improved sorting of unread/read/flagged messages in webmail",
		"IMP: Improved the behaviour of the Unsubscribe icon in Webmail (it no longer prompts with compose form, but rather presents confirmation)",
		"IMP: Improved the layout and content of the services overview screen when accessing the Synchonisation Service URL",
		"IMP: improved the layout of registration wizard",
		"IMP: Improved the resilience of Outlook 2013 ActiveSync synchronisation",
		"IMP: Improved the speed and resource usage when synchronising large mailboxes via EAS",
		"IMP: MailEnable no longer requires .Net 2/3.5 to be installed on systems that support .NET 4 (and later)",
		"IMP: MailEnable now adds Auto-Submitted: auto-replied header to autoresponders",
		"IMP: Message tracker application has been updated to make it consistent with the version in the MMC",
		"IMP: Migration Credential Capture now logs conversation and settings when validating to Config/Migrate/Debug folder",
		"IMP: Reduced the size of core Desktop WebMail web pages",
		"IMP: Significant ActiveSync caching and performance improvements ",
		"IMP: Significantly reduced memory footprint of IMAP service",
		"IMP: Speed of listing items in administration program improved",
		"IMP: Standard Edition no longer requires the installation of IIS Legacy Components",
		"IMP: Tracking utility did not show filter actions or MTA diagnostic information",
		"IMP: Updated Migration Console Documentation (to explain migration from legacy mail systems)",
		"IMP: Web security token files were being written to the root of the config directory (this has been moved to a more suitable location)",
		"IMP: WebMail client now tags cachable javacript files with the version number, so that client/cached versions will always be updated (removing erronous javascript scripting errors)",
		"IMP: Webmail layout and style improvements",
		"INF: Initial release of features as at http://www.mailenable.com/version7",
		"MailEnable Enterprise Premium Release Notes",
		"MOD: IMAP - Setting to tune the change behaviour of socket reads when processing command continuation requests (setting will allow devices sending partial requests accross different transmission units to be processed.",
		"MOD: LDAP address listing behavior reverted to list entries for each address rather than mailbox proxy addresses",
		"RCPT TO extended script handling for SMTP inbound was preserving recipient on failure",
		"VALUE NAME: Wait For Multiple Command Continuations",
		"VALUE: 1 (Enable, zero is off by default)",
	}

}
