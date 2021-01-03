on open droppedItems
	-- Change as you want
	set commandFlags to "-r -s -d"
	
	set commandPath to (quoted form of POSIX path of (path to me)) & "Contents/MacOS/mps"

	set posixPaths to {}
	
	repeat with thisItem in droppedItems
		set end of posixPaths to the quoted form of the POSIX path of thisItem
	end repeat
	
	set runCommand to commandPath & " " & commandFlags & " " & posixPaths
	tell application "Terminal"
		activate
		do script runCommand in front window
	end tell
end open