on open droppedItems
	-- Change to point to the mps binary
	set commandPath to "/path/to/mps"
	-- Change as you want
	set commandFlags to "-r -d"
	
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