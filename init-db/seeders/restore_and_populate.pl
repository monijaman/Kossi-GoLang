use strict;
use warnings;
use utf8;
use open qw(:std :utf8);

# Reverse Mappings (Bengali -> English)
# Keys are Regex patterns.
# We will sort them by length of the pattern string to prioritize longer matches.
my %rev_map = (
    'মেগাপিক্সেল' => 'MP',
    'পিক্সেল' => 'pixels',
    'গিগাবাইট' => 'GB',
    'টেরাবাইট' => 'TB',
    'এমএএইচ' => 'mAh',
    'ইঞ্চি' => 'inches',
    'মিমি' => 'mm',
    'গ্রাম' => 'g',
    'অ্যান্ড্রয়েড' => 'Android',
    'গ্লাস' => 'Glass',
    'প্লাস্টিক' => 'Plastic',
    'উপলব্ধ' => 'Available',
    'এলটিপিও' => 'LTPO',
    'অ্যামোলেড' => 'AMOLED',
    'এলসিডি' => 'LCD',
    'ওলেড' => 'OLED',
    'আইপিএস' => 'IPS',
    'নিট' => 'nits',
    'হার্টজ' => 'Hz',
    'সামনে' => 'front',
    'পিছনে' => 'back',
    'অ্যালুমিনিয়াম' => 'aluminum',
    'ফ্রেম' => 'frame',
    'নোনা-কোর' => 'Nona-core',
    'অক্টোবর' => 'October',
    'নভেম্বর' => 'November',
    'ডিসেম্বর' => 'December',
    'জানুয়ারি' => 'January',
    'ফেব্রুয়ারি' => 'February',
    'মার্চ' => 'March',
    'এপ্রিল' => 'April',
    'মে' => 'May',
    'জুন' => 'June',
    'জুলাই' => 'July',
    'আগস্ট' => 'August',
    'সেপ্টেম্বর' => 'September',
    '৫জি' => '5G',
    '৪জি' => '4G',
    '৩জি' => '3G',
    '২জি' => '2G',
    'জি' => 'G',
);

sub to_ascii_nums {
    my $s = shift;
    $s =~ tr/০-৯/0-9/;
    return $s;
}

sub to_english {
    my $s = shift;
    $s = to_ascii_nums($s);
    
    # Sort keys by length descending
    foreach my $pat (sort { length($b) <=> length($a) } keys %rev_map) {
        $s =~ s/$pat/$rev_map{$pat}/g;
    }
    return $s;
}

sub has_bengali {
    my $s = shift;
    return $s =~ /[\x{0980}-\x{09FF}]/;
}

my $file = shift or die "Usage: $0 <file>\n";
open my $in, "<:utf8", $file or die "Cannot open $file: $!";
my $content = do { local $/; <$in> };
close $in;

my %translations = ();

# 1. Parse existing getBanglaTranslations
if ($content =~ /(func.*getBanglaTranslations.*?return map\[string\]string\{)(.*?)(\})/s) {
    my $body = $2;
    # Use a more robust regex to capture keys and values
    while ($body =~ /"([^"]+)"\s*:\s*"([^"]+)"/g) {
        $translations{$1} = $2;
    }
}

# 2. Process specs map and collect new translations
if ($content =~ /(specs\s*:=\s*map\[string\]string\{)(.*?)(\})/s) {
    my $prefix = $1;
    my $body = $2;
    my $suffix = $3;
    
    my @lines = split /\n/, $body;
    my $new_body = "";
    
    foreach my $line (@lines) {
        if ($line =~ /"([^"]+)"\s*:\s*"([^"]+)"/) {
            my $k = $1;
            my $v = $2; 
            
            if (has_bengali($v)) {
                # Convert to English
                my $eng_v = to_english($v);
                
                # Add to translations map
                $translations{$eng_v} = $v;
                
                # Reconstruct line with English value
                if ($line =~ /^(\s*)/) { $new_body .= $1; }
                $new_body .= "\"$k\": \"$eng_v\",\n";
            } else {
                # Already English, keep as is
                $new_body .= "$line\n";
            }
        } else {
            $new_body .= "$line\n";
        }
    }
    chomp($new_body);
    $new_body .= "\n";
    
    $content =~ s/\Q$body\E/$new_body/s;
}

# 3. Rebuild getBanglaTranslations
if ($content =~ /(func.*getBanglaTranslations.*?return map\[string\]string\{)(.*?)(\})/s) {
    my $prefix = $1;
    my $suffix = $3;
    
    my $new_body = "\n";
    foreach my $eng (sort keys %translations) {
        $new_body .= "\t\t\"$eng\": \"$translations{$eng}\",\n";
    }
    
    $content =~ s/\Q$prefix$2$suffix\E/$prefix$new_body$suffix/s;
}

print $content;
