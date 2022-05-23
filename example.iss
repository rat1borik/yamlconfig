[Setup]
AppName=Example
AppVersion=1.0
DefaultDirName={commonpf}\Example

[Files]
Source: "yamlconfig.dll"; Flags: dontcopy
Source: "example.yaml"; DestDir: "{tmp}"; Flags: dontcopy

[Code]
type
  PWideChar = Cardinal;

function YAMLReadString(AFileName, APath, ADefault: AnsiString): PAnsiChar;
	external 'YAMLReadString@files:yamlconfig.dll cdecl';
function YAMLWriteString(AFileName, APath, AValue:AnsiString): Integer;
	external 'YAMLWriteString@files:yamlconfig.dll cdecl';
function MultiByteToWideChar(
    CodePage: UINT; dwFlags: DWORD; const lpMultiByteStr: AnsiString; cchMultiByte: Integer; 
    lpWideCharStr: string; cchWideChar: Integer): Integer;
  external 'MultiByteToWideChar@kernel32.dll stdcall';  

procedure AddToRTF(var Res: String; FuncName: String; Path: String; Value: AnsiString; Ok: Boolean);
begin
	if Ok then
		Res := Res + Format('{\i %s}: {\b %s} = {\cf1 %s}\line', [FuncName, Path, Value])
	else
		Res := Res + Format('{\i %s}: {\b %s} = {\cf2 %s}\line', [FuncName, Path, Value]);
	Res := Res + #13#10;
end;

var
	Page: TOutputMsgMemoWizardPage;

procedure InitializeWizard;
var
	rtf: String;
	fileName: AnsiString;
	strVal: AnsiString;
begin
	Page := CreateOutputMsgMemoPage(wpWelcome, 'Information', 'Display results', '', '');
	Page.RichEditViewer.UseRichEdit := True;

	ExtractTemporaryFile('example.yaml')
	fileName := ExpandConstant('{tmp}\example.yaml');

	rtf := '{\rtf1{\colortbl;\red0\green0\blue255;\red255\green0\blue0;}';

	strVal := YAMLReadString(fileName, 'foo.str', 'default');
	AddToRTF(rtf, 'YAMLReadString', 'foo.str', strVal, True);

	//rtf := rtf + '\line' + #13#10;

  strVal := YAMLReadString('жулик.txt', 'anything', 'default');
	AddToRTF(rtf, 'YAMLReadString', 'foo.str', strVal, True);

	rtf := rtf + '\line' + #13#10;

	// Write
	if YAMLWriteString(fileName, 'iss', 'InnoSetup')=0 {and YAMLReadString(fileName, 'foo.str', 'default', strVal, strLen)} then
		AddToRTF(rtf, 'YAMLWriteString', 'foo.str', {Copy(strVal, 1, strLen)}'ok', True)	
	else
		AddToRTF(rtf, 'YAMLWriteString', 'foo.str', 'failed', False);

	rtf := rtf + '}';

	Page.RichEditViewer.RTFText := rtf;
end;

function NextButtonClick(CurPageID: Integer): Boolean;
begin
	Result := not (CurPageID = Page.ID);
end;